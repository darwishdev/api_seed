package factory

import (
	"context"

	"github.com/meloneg/mln_data_pool/common/commontypes"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/xuri/excelize/v2"
)

type PublicFactoryInterface interface {
	SettingTypesBulkCreateParams() ([]string, error)
	SettingsBulkCreateParams(ctx context.Context, settingTypeFinder commontypes.IDFinder) ([]db.SettingsBulkCreateParams, error)
	IconsBulkCreateParams() ([]db.IconsBulkCreateParams, error)
}

type PublicFactory struct {
	data      *excelize.File
	separator string
}

func NewPublicFactory(data *excelize.File) PublicFactoryInterface {
	return &PublicFactory{
		data:      data,
		separator: ",",
	}
}
