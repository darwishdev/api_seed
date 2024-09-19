package repo

import (
	"context"
	"fmt"

	db "github.com/meloneg/mln_data_pool/common/db/gen"
)

func (r *PublicRepo) SettingTypesClear(ctx context.Context) error {
	err := r.store.SettingTypesClear(ctx)
	if err != nil {
		return fmt.Errorf("SettingTypesClear:%w", err)
	}
	return nil
}
func (r *PublicRepo) SettingTypesBulkCreate(ctx context.Context, arg []string) (int64, error) {
	count, err := r.store.SettingTypesBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("SettingTypesBulkCreate: %w", err)
	}
	return count, nil
}

func (r *PublicRepo) IconsBulkCreate(ctx context.Context, arg []db.IconsBulkCreateParams) (int64, error) {
	count, err := r.store.IconsBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("IconsBulkCreate: %w", err)
	}
	return count, nil
}
func (r *PublicRepo) SettingTypeFindByType(ctx context.Context, arg string) (int32, error) {

	cityId, err := r.store.SettingTypeFindByType(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("SettingTypeFindByType: %w", err)
	}
	return cityId, nil
}

func (r *PublicRepo) SettingsClear(ctx context.Context) error {
	err := r.store.SettingsClear(ctx)
	if err != nil {
		return fmt.Errorf("SettingsClear:%w", err)
	}
	return nil
}
func (r *PublicRepo) SettingsBulkCreate(ctx context.Context, arg []db.SettingsBulkCreateParams) (int64, error) {
	count, err := r.store.SettingsBulkCreate(context.Background(), arg)
	if err != nil {
		return 0, fmt.Errorf("SettingsBulkCreate: %w", err)
	}
	return count, nil
}
