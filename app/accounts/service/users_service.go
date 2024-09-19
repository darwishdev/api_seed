package service

import (
	"context"
	"fmt"

	"github.com/meloneg/mln_data_pool/app/accounts/types"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supabase"
	"github.com/rs/zerolog/log"
)

func (s *AccountService) PermissionsBulkCreate(c context.Context) (permissions []db.PermissionsBulkCreateParams, err error) {
	permissions, err = s.accountsFactory.PermissionsBulkCreateParams()
	if err != nil {
		return nil, err
	}

	err = s.repo.PermissionsClear(c)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.PermissionsBulkCreate(c, permissions)
	if err != nil {
		return nil, err
	}
	return
}

func (s *AccountService) NavigationBarsBulkCreate(c context.Context) (items []db.NavigationBarCreateParams, err error) {
	items, err = s.accountsFactory.NavigationBarsBulkCreateParams(c, s.repo.IconFindByName)
	if err != nil {
		return nil, err
	}

	for _, v := range items {
		err = s.repo.NavigationBarCreate(c, v)
		if err != nil {
			return nil, err
		}

	}
	return
}

func (s *AccountService) loadPermissionIds(functionNames []string, resp *[]int32) error {
	permissionsIds, err := s.repo.PermissionIdsByFunctions(context.Background(), functionNames)
	if err != nil {
		return fmt.Errorf("PermissionIdsByFunctions: %w", err)
	}

	log.Debug().Interface("response", permissionsIds).Msg("loadPermissionIdsloadPermissionIds")
	*resp = permissionsIds

	return nil
}

func (s *AccountService) loadRoleIds(roleNames []string, resp *[]int32) error {
	roleIds, err := s.repo.RoleIdsByNames(context.Background(), roleNames)
	if err != nil {
		return fmt.Errorf("loadRoleIds: %w", err)
	}
	*resp = roleIds

	return nil
}

func (s *AccountService) loadAllPermissionIds(count int, resp *[]int32) {
	for i := 0; i < count; i++ {
		*resp = append(*resp, int32(i+1))
	}
}

func (s *AccountService) RolesBulkCreate(c context.Context) (roles []types.RolesBulkCreateResponse, err error) {
	rows, err := s.data.GetRows("roles")
	if err != nil {
		return nil, err
	}
	err = s.repo.RolesClear(c)
	if err != nil {
		return nil, err
	}
	return s.RolesBulkCreateBase(c, rows, 0)
}
func (s *AccountService) RolesBulkCreateFile(c context.Context, rows [][]string) (roles []types.RolesBulkCreateResponse, err error) {
	rolesCount, err := s.repo.RolesCount(c)
	if err != nil {
		return nil, err
	}
	return s.RolesBulkCreateBase(c, rows, rolesCount)
}

func (s *AccountService) RolesBulkCreateBase(c context.Context, rows [][]string, rolesCount int64) (roles []types.RolesBulkCreateResponse, err error) {
	params, err := s.accountsFactory.RolesBulkCreateParams(rows, rolesCount)
	if err != nil {
		return nil, err
	}
	var roleParams []db.RolesBulkCreateParams
	var permissions []db.RolePermissionsBulkCreateParams
	for _, resp := range *params {
		roleParams = append(roleParams, resp.Role)
		roles = append(roles, resp)
		permissionsIds := make([]int32, 0)
		if resp.RolePermissions[0] == "*" {
			s.loadAllPermissionIds(resp.PermissionsLen, &permissionsIds)
		} else {
			s.loadPermissionIds(resp.RolePermissions, &permissionsIds)
		}

		var currenRolePermissions []db.RolePermissionsBulkCreateParams
		for _, v := range permissionsIds {
			permission := db.RolePermissionsBulkCreateParams{
				RoleID:       int32(resp.RoleID),
				PermissionID: int32(v),
			}
			currenRolePermissions = append(currenRolePermissions, permission)
		}

		permissions = append(permissions, currenRolePermissions...)
	}

	_, err = s.repo.RolesBulkCreate(c, roleParams)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.RolePermissionsBulkCreate(c, permissions)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
func (s *AccountService) UsersBulkCreate(ctx context.Context) ([]types.UsersBulkCreateResponse, error) {
	params, err := s.accountsFactory.UsersBulkCreateParams(ctx)

	log.Debug().Interface("o", params).Msg("l")
	if err != nil {
		return nil, err
	}
	var userParams []db.UsersBulkCreateParams
	var users []types.UsersBulkCreateResponse
	var rolesParams []db.UserRolesBulkCreateParams
	for _, resp := range *params {

		userParams = append(userParams, resp.User)
		if err != nil {
			return nil, err
		}
		users = append(users, resp)
		roleIds := make([]int32, 0)
		s.loadRoleIds(resp.UserRoles, &roleIds)

		var currenUserRoles []db.UserRolesBulkCreateParams
		for _, v := range roleIds {
			role := db.UserRolesBulkCreateParams{
				UserID: int32(resp.UserID),
				RoleID: v,
			}
			currenUserRoles = append(currenUserRoles, role)
		}

		rolesParams = append(rolesParams, currenUserRoles...)
	}

	_, err = s.repo.UsersBulkCreate(ctx, userParams)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.UserRolesBulkCreate(ctx, rolesParams)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *AccountService) CustomersBulkCreate(ctx context.Context) ([]types.CustomersBulkCreateResponse, error) {
	params, err := s.accountsFactory.CustomersBulkCreateParams(ctx)

	log.Debug().Interface("o", params).Msg("l")
	if err != nil {
		return nil, err
	}
	var customerParams []db.CustomersBulkCreateParams
	for _, resp := range *params {

		customerParams = append(customerParams, resp.Customer)

	}
	_, err = s.repo.CustomersBulkCreate(ctx, customerParams)
	if err != nil {
		return nil, err
	}

	return *params, nil
}
func (s *AccountService) OwnersBulkCreate(ctx context.Context) (*types.OwnersBulkCreateResponse, error) {
	params, err := s.accountsFactory.OwnersBulkCreateParams(ctx)

	if err != nil {
		return nil, err
	}

	_, err = s.repo.OwnersBulkCreate(ctx, params.Owners)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (s *AccountService) AuthUsersBulkCreate(ctx context.Context) ([]supabase.AuthenticatedDetails, error) {
	// params, err := s.accountsFactory.UsersBulkCreateParams()
	// if err != nil {
	// 	return nil, err
	// }
	// var users []supabase.AuthenticatedDetails
	// for _, resp := range *params {
	// 	user, err := s.supa.SignUp(ctx, resp.AuthUser)
	// 	time.Sleep(300 * time.Microsecond)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, *user)

	// }

	return nil, nil
}
