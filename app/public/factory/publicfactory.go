package factory

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/meloneg/mln_data_pool/common/commontypes"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/rs/zerolog/log"
)

// supa "github.com/darwishdev/supabase-go"

func (f *PublicFactory) SettingTypesBulkCreateParams() ([]string, error) {
	rows, err := f.data.GetRows("Settings")
	if err != nil {
		return nil, err
	}
	resp := make([]string, 0)
	for rowIndex, row := range rows {
		// Skip the header row (if needed)
		if rowIndex == 0 {
			continue
		}

		if err != nil {
			return nil, err
		}
		if len(row) > 0 {
			if row[0] != "" {
				resp = append(resp, row[0])
			}

		}
	}
	return resp, nil
}

func convertFilename(filename string) string {
	// parts := strings.SplitN(filename, "-", 2)
	// if len(parts) < 2 {
	// 	return ""
	// }

	// // Remove everything before the first hyphen (including the hyphen)
	// description := parts[1]

	// // Remove the .svg extension
	// description = strings.TrimSuffix(description, ".svg")

	// // Replace spaces with hyphens
	// convertedName := strings.ReplaceAll(description, " ", "-")
	return strings.TrimSuffix(filename, ".svg")
}

func (f *PublicFactory) IconsBulkCreateParams() ([]db.IconsBulkCreateParams, error) {
	directory := "icons"
	// Read all files in the directory
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	resp := make([]db.IconsBulkCreateParams, 0)
	resultMap := make(map[string]bool, 0)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".svg" {
			cwd, err := os.Getwd()
			if err != nil {
				return nil, err
			}
			// Prepend the current directory path to the filename
			fullPath := filepath.Join(cwd, directory, file.Name())

			content, err := os.ReadFile(fullPath)

			if err != nil {
				return nil, err
			}

			convertedName := convertFilename(file.Name())

			_, ok := resultMap[convertedName]
			if !ok {
				resultMap[convertedName] = true

				resp = append(resp, db.IconsBulkCreateParams{
					IconName:    convertedName,
					IconContent: string(content),
				})
			}

			// convertedName = convertedName + ".svg"

			// err = os.Rename(fullPath, filepath.Join(cwd, directory, convertedName))
			// if err != nil {
			// 	return nil, err
			// }
		}
	}

	return resp, nil
}

func (f *PublicFactory) SettingsBulkCreateParams(ctx context.Context, settingTypeFinder commontypes.IDFinder) ([]db.SettingsBulkCreateParams, error) {
	rows, err := f.data.GetRows("Settings")
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("recod", len(rows)).Msg("rpwsss")
	resp := make([]db.SettingsBulkCreateParams, 0)
	for rowIndex, row := range rows {
		log.Debug().Interface("recod", len(rows)).Msg("rpwsss")
		// Skip the header row (if needed)
		if rowIndex == 0 {
			continue
		}

		if err != nil {
			return nil, err
		}
		if len(row) > 2 {
			if row[2] != "" {
				typeId, err := settingTypeFinder(ctx, row[2])
				if err != nil {
					return nil, err
				}
				recod := db.SettingsBulkCreateParams{
					SettingTypeID: typeId,
					SettingKey:    row[3],
					SettingValue:  row[4],
				}

				resp = append(resp, recod)
			}

		}
	}
	return resp, nil
}
