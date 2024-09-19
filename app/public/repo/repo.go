package repo

import (
	"context"

	db "github.com/meloneg/mln_data_pool/common/db/gen"
)

type PublicRepoInterface interface {
	SettingTypesClear(ctx context.Context) error
	SettingTypesBulkCreate(ctx context.Context, arg []string) (int64, error)
	SettingTypeFindByType(ctx context.Context, arg string) (int32, error)
	SettingsClear(ctx context.Context) error
	SettingsBulkCreate(ctx context.Context, arg []db.SettingsBulkCreateParams) (int64, error)
	IconsBulkCreate(ctx context.Context, arg []db.IconsBulkCreateParams) (int64, error)
}

type PublicRepo struct {
	store db.Store
}

func NewPublicRepo(store db.Store) PublicRepoInterface {
	return &PublicRepo{
		store: store,
	}
}
