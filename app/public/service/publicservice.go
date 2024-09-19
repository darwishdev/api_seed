package service

import (
	"context"

	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/rs/zerolog/log"
)

func (s *PublicService) SettingTypesBulkCreate(c context.Context) ([]string, error) {
	settingTypes, err := s.publicFactory.SettingTypesBulkCreateParams()
	if err != nil {
		return nil, err
	}

	_, err = s.repo.SettingTypesBulkCreate(c, settingTypes)
	if err != nil {
		return nil, err
	}
	return settingTypes, nil
}

func (s *PublicService) SettingsBulkCreate(c context.Context) ([]db.SettingsBulkCreateParams, error) {
	settings, err := s.publicFactory.SettingsBulkCreateParams(c, s.repo.SettingTypeFindByType)
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("setttings", settings).Msg("SettingsBulkCreate")

	_, err = s.repo.SettingsBulkCreate(c, settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (s *PublicService) IconsBulkCreate(c context.Context) ([]db.IconsBulkCreateParams, error) {
	icons, err := s.publicFactory.IconsBulkCreateParams()
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("setttings", icons).Msg("SettingsBulkCreate")

	_, err = s.repo.IconsBulkCreate(c, icons)
	if err != nil {
		return nil, err
	}
	return icons, nil
}
