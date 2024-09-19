package factory

import (
	"context"

	"github.com/meloneg/mln_data_pool/app/accounts/types"

	"github.com/meloneg/mln_data_pool/common/commontypes"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supaclient"
	"github.com/xuri/excelize/v2"
)

type AccountsFactoryInterface interface {
	PermissionsBulkCreateParams() ([]db.PermissionsBulkCreateParams, error)
	NavigationBarsBulkCreateParams(ctx context.Context, iconFinder commontypes.IDFinder) ([]db.NavigationBarCreateParams, error)
	RolesBulkCreateParams(rows [][]string, rolesCount int64) (*[]types.RolesBulkCreateResponse, error)
	OwnersBulkCreateParams(ctx context.Context) (*types.OwnersBulkCreateResponse, error)
	UsersBulkCreateParams(ctx context.Context) (*[]types.UsersBulkCreateResponse, error)
	CustomersBulkCreateParams(ctx context.Context) (*[]types.CustomersBulkCreateResponse, error)
}

type AccountsFactory struct {
	data      *excelize.File
	supa      supaclient.SupabaseServiceInterface
	separator string
}

func NewAccountsFactory(data *excelize.File, supa supaclient.SupabaseServiceInterface) AccountsFactoryInterface {
	return &AccountsFactory{
		supa:      supa,
		data:      data,
		separator: ",",
	}
}
