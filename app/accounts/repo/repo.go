package repo

import (
	"context"

	db "github.com/meloneg/mln_data_pool/common/db/gen"
)

type AccountRepoInterface interface {
	PermissionsClear(ctx context.Context) error
	NavigationBarsClear(ctx context.Context) error
	PermissionsBulkCreate(ctx context.Context, arg []db.PermissionsBulkCreateParams) (int64, error)
	NavigationBarCreate(ctx context.Context, arg db.NavigationBarCreateParams) error
	NavigationBarsBulkCreateParamsBulkCreate(ctx context.Context, arg []db.PermissionsBulkCreateParams) (int64, error)
	PermissionIdsByFunctions(ctx context.Context, functionsNames []string) ([]int32, error)
	RoleIdsByNames(ctx context.Context, rolesNames []string) ([]int32, error)
	RolesClear(ctx context.Context) error
	RolesCount(ctx context.Context) (int64, error)
	RolesBulkCreate(ctx context.Context, arg []db.RolesBulkCreateParams) (int64, error)
	RolePermissionsBulkCreate(ctx context.Context, arg []db.RolePermissionsBulkCreateParams) (int64, error)
	UsersClear(ctx context.Context) error
	UsersBulkCreate(ctx context.Context, arg []db.UsersBulkCreateParams) (int64, error)
	IconFindByName(ctx context.Context, name string) (int32, error)
	OwnersBulkCreate(ctx context.Context, arg []db.OwnersBulkCreateParams) (int64, error)
	UserRolesBulkCreate(ctx context.Context, arg []db.UserRolesBulkCreateParams) (int64, error)
	AuthUserFindIdByEmail(ctx context.Context, email string) (string, error)
	CustomersBulkCreate(ctx context.Context, arg []db.CustomersBulkCreateParams) (int64, error)
}

type AccountRepo struct {
	store db.Store
}

func NewAccountRepo(store db.Store) AccountRepoInterface {
	return &AccountRepo{
		store: store,
	}
}
