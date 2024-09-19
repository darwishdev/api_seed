package factory

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/meloneg/mln_data_pool/app/accounts/types"
	"github.com/meloneg/mln_data_pool/common/commontypes"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supabase"
	"github.com/rs/zerolog/log"

	// supa "github.com/darwishdev/supabase-go"
	"golang.org/x/crypto/bcrypt"
)

func (f *AccountsFactory) PermissionsBulkCreateParams() ([]db.PermissionsBulkCreateParams, error) {
	rows, err := f.data.GetRows("permissions")
	if err != nil {
		return nil, err
	}
	resp := make([]db.PermissionsBulkCreateParams, 0)
	for rowIndex, row := range rows {
		// Skip the header row (if needed)
		if rowIndex == 0 {
			continue
		}

		if err != nil {

			return nil, err
		}
		record := db.PermissionsBulkCreateParams{
			PermissionName:        row[0],
			PermissionFunction:    row[1],
			PermissionDescription: pgtype.Text{String: row[2], Valid: true},
			PermissionGroup:       row[3],
		}

		resp = append(resp, record)
	}

	return resp, nil
}

func (f *AccountsFactory) NavigationBarsBulkCreateParams(ctx context.Context, iconFinder commontypes.IDFinder) ([]db.NavigationBarCreateParams, error) {
	rows, err := f.data.GetRows("navigation_bars")
	if err != nil {
		return nil, err
	}
	resp := make([]db.NavigationBarCreateParams, 0)
	for rowIndex, row := range rows {

		// Skip the header row (if needed)
		if rowIndex == 0 {
			continue
		}
		if row[0] == "" {
			break
		}

		if err != nil {
			return nil, err
		}

		hasTo := (row[3] != "" && row[3] != "NULL" && row[3] != "null")
		hasParent := (row[4] != "" && row[4] != "NULL" && row[4] != "null")

		iconId, err := iconFinder(ctx, row[2])

		if err != nil {
			return nil, err
		}

		record := db.NavigationBarCreateParams{
			MenuKey:        row[0],
			Label:          row[1],
			IconID:         iconId,
			Route:          pgtype.Text{String: row[3], Valid: hasTo},
			ParentKey:      pgtype.Text{String: row[4], Valid: hasParent},
			PermissionName: row[5],
			LabelAr:        pgtype.Text{String: row[6], Valid: true},
		}

		resp = append(resp, record)
	}

	return resp, nil
}

func (f *AccountsFactory) RolesBulkCreateParams(rows [][]string, rolesCount int64) (*[]types.RolesBulkCreateResponse, error) {

	permissionsRows, err := f.data.GetRows("permissions")
	if err != nil {
		return nil, err
	}

	response := make([]types.RolesBulkCreateResponse, 0)
	for rowIndex, row := range rows {
		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		dbRecord := db.RolesBulkCreateParams{
			RoleName:        row[0],
			RoleDescription: pgtype.Text{String: row[1], Valid: true},
		}
		currenPermissions := strings.Split(row[2], f.separator)
		for i := 0; i < len(currenPermissions); i++ {
			currenPermissions[i] = strings.TrimSpace(currenPermissions[i])
		}

		var permissionsLen int = len(currenPermissions)
		if currenPermissions[0] == "*" {
			permissionsLen = len(permissionsRows) - 1
		}

		record := types.RolesBulkCreateResponse{
			Role:            dbRecord,
			RoleID:          int(rolesCount) + rowIndex,
			RolePermissions: currenPermissions,
			PermissionsLen:  permissionsLen, // here we substract 1 to remove the header
		}

		response = append(response, record)

	}

	return &response, nil
}

func (f *AccountsFactory) UsersBulkCreateParams(ctx context.Context) (*[]types.UsersBulkCreateResponse, error) {
	rows, err := f.data.GetRows("users")
	if err != nil {
		return nil, err
	}
	response := make([]types.UsersBulkCreateResponse, 0)
	for rowIndex, row := range rows {

		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		email := strings.TrimSpace(row[1])
		authUserReq := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		_, err := f.supa.SignUp(ctx, authUserReq)
		if err != nil {
			log.Debug().Str("cannot insert the auth user for", authUserReq.Email).Err(err).Msg("auth debug")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(row[4]), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		dbRecord := db.UsersBulkCreateParams{
			UserName:     row[0],
			UserEmail:    email,
			UserImage:    pgtype.Text{String: row[2], Valid: true},
			UserPhone:    pgtype.Text{String: row[3], Valid: true},
			UserPassword: string(hashedPassword),
			UserCode:     row[6],
			// AccountID:    pgtype.Int4{Int32: accountId, Valid: accountId != 0},
		}

		authUser := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		currenRoles := strings.Split(row[5], f.separator)
		for i := 0; i < len(currenRoles); i++ {
			currenRoles[i] = strings.TrimSpace(currenRoles[i])
		}

		record := types.UsersBulkCreateResponse{
			User:      dbRecord,
			AuthUser:  authUser,
			UserID:    rowIndex,
			UserRoles: currenRoles,
			RolesLen:  len(currenRoles),
		}
		response = append(response, record)

	}

	return &response, nil
}

func (f *AccountsFactory) CustomersBulkCreateParams(ctx context.Context) (*[]types.CustomersBulkCreateResponse, error) {
	rows, err := f.data.GetRows("customers")
	if err != nil {
		return nil, err
	}
	response := make([]types.CustomersBulkCreateResponse, 0)
	for rowIndex, row := range rows {

		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		email := strings.TrimSpace(row[1])
		authUserReq := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		duration := 3 * time.Second
		time.Sleep(duration)
		_, err := f.supa.SignUp(ctx, authUserReq)
		if err != nil {
			log.Debug().Str("cannot insert the auth user for", authUserReq.Email).Err(err).Msg("auth debug")
			// return nil, fmt.Errorf("failed to insert auth: %w", err)

		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(row[4]), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		dbRecord := db.CustomersBulkCreateParams{
			CustomerName:       row[0],
			CustomerEmail:      email,
			CustomerImage:      pgtype.Text{String: row[2], Valid: true},
			CustomerPhone:      pgtype.Text{String: row[3], Valid: true},
			CustomerPassword:   string(hashedPassword),
			CustomerCode:       row[5],
			CustomerNationalID: pgtype.Text{String: row[6], Valid: true},
			// AccountID:    pgtype.Int4{Int32: accountId, Valid: accountId != 0},
		}

		authUser := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		currenRoles := strings.Split(row[5], f.separator)
		for i := 0; i < len(currenRoles); i++ {
			currenRoles[i] = strings.TrimSpace(currenRoles[i])
		}

		record := types.CustomersBulkCreateResponse{
			Customer:   dbRecord,
			AuthUser:   authUser,
			CustomerID: rowIndex,
		}
		response = append(response, record)

	}

	return &response, nil
}

func (f *AccountsFactory) OwnersBulkCreateParams(ctx context.Context) (*types.OwnersBulkCreateResponse, error) {
	rows, err := f.data.GetRows("owners")
	if err != nil {
		return nil, err
	}

	records := make([]db.OwnersBulkCreateParams, 0)
	authRecords := make([]supabase.UserCredentials, 0)
	for rowIndex, row := range rows {

		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		email := strings.TrimSpace(row[1])
		authUserReq := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		_, err := f.supa.SignUp(ctx, authUserReq)
		if err != nil {
			log.Debug().Str("cannot insert the auth user for", authUserReq.Email).Err(err).Msg("auth debug")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(row[4]), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		dbRecord := db.OwnersBulkCreateParams{
			OwnerName:     row[0],
			OwnerEmail:    email,
			OwnerImage:    pgtype.Text{String: row[2], Valid: true},
			OwnerPhone:    pgtype.Text{String: row[3], Valid: true},
			OwnerPassword: string(hashedPassword),
			// AccountID:    pgtype.Int4{Int32: accountId, Valid: accountId != 0},
			OwnerNationalID: row[5],
		}

		authOwner := supabase.UserCredentials{
			Email:    email,
			Password: row[4],
		}
		currenRoles := strings.Split(row[5], f.separator)
		for i := 0; i < len(currenRoles); i++ {
			currenRoles[i] = strings.TrimSpace(currenRoles[i])
		}

		records = append(records, dbRecord)
		authRecords = append(authRecords, authOwner)

	}

	return &types.OwnersBulkCreateResponse{
		AuthUsers: authRecords,
		Owners:    records,
	}, nil
}
