package file

import (
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
)

func LoadFile(path string) (*excelize.File, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Debug().Err(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Debug().Err(err)

		}
	}()

	return f, nil

}
