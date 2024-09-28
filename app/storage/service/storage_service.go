package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/meloneg/mln_data_pool/supabase"
	"github.com/meloneg/mln_data_pool/supaclient"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// var publicImageBucketName = "image"
var (
	assetsFolderName = "assets"
	folders          = []string{"images", "properties", "units", "rooms"}
)

func parseImageFile(path string) *bytes.Buffer {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil
	}
	imageData := &bytes.Buffer{}
	_, err = io.Copy(imageData, imageFile)
	if err != nil {
		return nil
	}
	defer imageFile.Close()
	return imageData
}

func (s *StorageService) BucketCreateImages(c context.Context) error {
	for _, folder := range folders {
		req := supabase.BucketOption{
			Id:     folder,
			Name:   folder,
			Public: true,
		}
		_, err := s.supa.BucketCreate(c, req)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageService) UploadInitialImages(c context.Context) error {
	for _, folder := range folders {
		directory := assetsFolderName + "/" + folder
		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				extension := strings.ToLower(filepath.Ext(path))
				if extension == ".jpg" || extension == ".jpeg" || extension == ".png" || extension == ".webp" {
					imageData := parseImageFile(path)
					req := supaclient.BucketUploadRequest{
						BucketName: folder,
						Path:       "./" + info.Name(),
						Reader:     imageData,
						FileType:   "image/" + extension[1:],
					}

					_, err = s.supa.BucketUpload(req)

					if err != nil {
						// Log the error but don't stop the execution
						fmt.Printf("Failed to upload %s: %v\n", info.Name(), err)
					}
				}
			}

			return nil
		})
		if err != nil {
			return nil
		}
	}
	return nil
}
