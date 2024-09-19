package repo

import (
	"context"
	"fmt"

	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/rs/zerolog/log"
)

func (r *AccountRepo) PermissionsClear(ctx context.Context) error {

	err := r.store.PermissionsClear(ctx)
	if err != nil {
		return fmt.Errorf("error happend :%w", err)
	}
	return nil

}

func (r *AccountRepo) NavigationBarsClear(ctx context.Context) error {

	err := r.store.NavigationBarsClear(ctx)
	if err != nil {
		return fmt.Errorf("error happend :%w", err)
	}
	return nil

}

func (r *AccountRepo) PermissionsBulkCreate(ctx context.Context, arg []db.PermissionsBulkCreateParams) (int64, error) {
	log.Debug().Interface("arg", arg).Msg("asdasdasdasd")

	count, err := r.store.PermissionsBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("err %w", err)
	}
	return count, nil
}
func (r *AccountRepo) NavigationBarCreate(ctx context.Context, arg db.NavigationBarCreateParams) error {
	err := r.store.NavigationBarCreate(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("failed to get navigations: %w", err)
	}
	return nil
}

func (r *AccountRepo) NavigationBarsBulkCreateParamsBulkCreate(ctx context.Context, arg []db.PermissionsBulkCreateParams) (int64, error) {

	count, err := r.store.PermissionsBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("PermissionsBulkCreate: %w", err)
	}
	return count, nil
}
func (r *AccountRepo) PermissionIdsByFunctions(ctx context.Context, functionsNames []string) ([]int32, error) {

	permissionsIds, err := r.store.PermissionIdsByFunctions(context.Background(), functionsNames)
	if err != nil {
		return nil, fmt.Errorf("PermissionIdsByFunctions: %w", err)
	}
	return permissionsIds, nil
}

func (r *AccountRepo) RoleIdsByNames(ctx context.Context, rolesNames []string) ([]int32, error) {

	rolesIds, err := r.store.RoleIdsByNames(context.Background(), rolesNames)
	if err != nil {
		return nil, fmt.Errorf("RoleIdsByNames: %w", err)
	}
	return rolesIds, nil
}
func (r *AccountRepo) RolesClear(ctx context.Context) error {

	err := r.store.RolesClear(ctx)
	if err != nil {
		return fmt.Errorf("RolesClear: %w", err)
	}

	return nil
}
func (r *AccountRepo) RolesCount(ctx context.Context) (int64, error) {

	count, err := r.store.RolesCount(context.Background())
	if err != nil {
		return 0, fmt.Errorf("RolesCount: %w", err)
	}
	return count, nil

}

func (r *AccountRepo) RolesBulkCreate(ctx context.Context, arg []db.RolesBulkCreateParams) (int64, error) {

	count, err := r.store.RolesBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("RolesBulkCreate: %w", err)
	}
	return count, nil

}
func (r *AccountRepo) RolePermissionsBulkCreate(ctx context.Context, arg []db.RolePermissionsBulkCreateParams) (int64, error) {

	log.Debug().Interface("repo permissions", arg).Msg("RolePermissionsBulkCreate")

	count, err := r.store.RolePermissionsBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("RolePermissionsBulkCreate: %w", err)
	}
	return count, nil

}

func (r *AccountRepo) UsersClear(ctx context.Context) error {

	err := r.store.UsersClear(ctx)
	if err != nil {
		return fmt.Errorf("UsersClear : %w", err)
	}

	return nil
}
func (r *AccountRepo) UsersBulkCreate(ctx context.Context, arg []db.UsersBulkCreateParams) (int64, error) {
	count, err := r.store.UsersBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("UsersBulkCreate: %w", err)
	}
	return count, nil

}

func (r *AccountRepo) CustomersBulkCreate(ctx context.Context, arg []db.CustomersBulkCreateParams) (int64, error) {
	count, err := r.store.CustomersBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("CustomersBulkCreate: %w", err)
	}

	log.Debug().Interface("Asdasdadasd", count).Interface("asdasdrggg", arg).Msg("CustomersBulkCreate")
	return count, nil

}

func (r *AccountRepo) OwnersBulkCreate(ctx context.Context, arg []db.OwnersBulkCreateParams) (int64, error) {
	count, err := r.store.OwnersBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("OwnersBulkCreate: %w", err)
	}
	return count, nil

}
func (r *AccountRepo) UserRolesBulkCreate(ctx context.Context, arg []db.UserRolesBulkCreateParams) (int64, error) {

	count, err := r.store.UserRolesBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("failed to get permission ids: %w", err)
	}
	return count, nil

}

func (r *AccountRepo) AuthUserFindIdByEmail(ctx context.Context, email string) (string, error) {
	var id string
	err := r.store.QueryRow(ctx, "SELECT id::text FROM auth.users WHERE email = $1 LIMIT 1", email).Scan(&id)
	log.Debug().Interface("id", id).Err(err).Msg("hello")
	if err != nil {
		return "", err
	}
	return "", nil

}
func (r *AccountRepo) IconFindByName(ctx context.Context, name string) (int32, error) {
	id, err := r.store.IconFindByName(context.Background(), name)
	if err != nil {
		return 0, fmt.Errorf("IconFindByName: %w", err)
	}
	return id, nil

}
