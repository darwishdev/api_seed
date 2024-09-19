package service

import (
	"context"

	"github.com/meloneg/mln_data_pool/app/public/factory"
	"github.com/meloneg/mln_data_pool/app/public/repo"

	// supa "github.com/darwishdev/supabase-go"
	"github.com/xuri/excelize/v2"
)

type PublicServiceInterface interface {
	SettingTypesBulkCreate(c context.Context) ([]string, error)
	SettingsBulkCreate(c context.Context) ([]string, error)
}

type PublicService struct {
	data          *excelize.File
	publicFactory factory.PublicFactoryInterface
	repo          repo.PublicRepoInterface
}

func NewPublicService(data *excelize.File, publicFactory factory.PublicFactoryInterface, publicRepo repo.PublicRepoInterface) *PublicService {
	return &PublicService{
		data:          data,
		publicFactory: publicFactory,
		repo:          publicRepo,
	}
}
