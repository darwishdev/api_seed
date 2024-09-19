package routes

import (
	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Test bool `json:"test"`
}

func (r *PublicRoutes) IconsBulkCreate(c *fiber.Ctx) error {
	paymentTypes, err := r.service.IconsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(paymentTypes)
}

func (r *PublicRoutes) SettingsBulkCreate(c *fiber.Ctx) error {
	paymentTypes, err := r.service.SettingsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(paymentTypes)
}

func (r *PublicRoutes) PublicSchemaSeed(c *fiber.Ctx) error {
	settingTypes, err := r.service.SettingTypesBulkCreate(c.Context())
	if err != nil {
		return err
	}
	settings, err := r.service.SettingsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	response := map[string]any{
		"settingTypes": settingTypes,
		"settings":     settings,
	}
	return c.JSON(response)
}
func (r *PublicRoutes) SettingTypesBulkCreate(c *fiber.Ctx) error {
	settings, err := r.service.SettingTypesBulkCreate(c.Context())
	if err != nil {
		return err
	}
	response := map[string]any{
		"settings": settings,
	}
	return c.JSON(response)
}
