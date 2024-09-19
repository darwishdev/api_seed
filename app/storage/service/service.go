package service

import (
	"context"

	"github.com/meloneg/mln_data_pool/supaclient"
)

type StorageServiceInterface interface {
	BucketCreateImages(c context.Context) error
	UploadInitialImages(c context.Context) error
}

type StorageService struct {
	supa supaclient.SupabaseServiceInterface
}

func NewStorageService(supa supaclient.SupabaseServiceInterface) *StorageService {
	return &StorageService{
		supa: supa,
	}
}
