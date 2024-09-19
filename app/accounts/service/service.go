package service

import (
	"context"

	"github.com/meloneg/mln_data_pool/app/accounts/factory"
	"github.com/meloneg/mln_data_pool/app/accounts/repo"
	"github.com/meloneg/mln_data_pool/app/accounts/types"

	// entitiesRepo "github.com/meloneg/mln_data_pool/app/entities/repo"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supabase"
	"github.com/meloneg/mln_data_pool/supaclient"

	// supa "github.com/darwishdev/supabase-go"
	"github.com/xuri/excelize/v2"
)

type AccountServiceInterface interface {
	PermissionsBulkCreate(c context.Context) (permissions []db.PermissionsBulkCreateParams, err error)
	NavigationBarsBulkCreate(c context.Context) (items []db.NavigationBarCreateParams, err error)
	loadPermissionIds(functionNames []string, resp *[]int32) error
	loadRoleIds(roleNames []string, resp *[]int32) error
	loadAllPermissionIds(count int, resp *[]int32)
	RolesBulkCreate(c context.Context) (roles []types.RolesBulkCreateResponse, err error)
	RolesBulkCreateFile(c context.Context, rows [][]string) (roles []types.RolesBulkCreateResponse, err error)
	RolesBulkCreateBase(c context.Context, rows [][]string, rolesCount int64) (roles []types.RolesBulkCreateResponse, err error)
	UsersBulkCreate(ctx context.Context) (users []types.UsersBulkCreateResponse, err error)
	CustomersBulkCreate(ctx context.Context) ([]types.CustomersBulkCreateResponse, error)
	OwnersBulkCreate(ctx context.Context) (*types.OwnersBulkCreateResponse, error)
	AuthUsersBulkCreate(ctx context.Context) ([]supabase.AuthenticatedDetails, error)
}

type AccountService struct {
	data            *excelize.File
	accountsFactory factory.AccountsFactoryInterface
	repo            repo.AccountRepoInterface
	supa            supaclient.SupabaseServiceInterface
}

func NewAccountService(data *excelize.File, accountsFactory factory.AccountsFactoryInterface, accountRepo repo.AccountRepoInterface, supa supaclient.SupabaseServiceInterface) *AccountService {
	return &AccountService{
		data:            data,
		supa:            supa,
		accountsFactory: accountsFactory,
		repo:            accountRepo,
	}
}
