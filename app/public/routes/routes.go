package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meloneg/mln_data_pool/app/public/factory"
	"github.com/meloneg/mln_data_pool/app/public/repo"
	"github.com/meloneg/mln_data_pool/app/public/service"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/xuri/excelize/v2"
)

type PublicRoutesInterface interface {
	RegisterPublicRoutes(app *fiber.App)
}

type PublicRoutes struct {
	service service.PublicService
}

func NewPublicRoutes(store db.Store, data *excelize.File) PublicRoutesInterface {
	publicFactory := factory.NewPublicFactory(data)
	publicRepo := repo.NewPublicRepo(store)
	publicService := service.NewPublicService(data, publicFactory, publicRepo)
	return &PublicRoutes{
		service: *publicService,
	}
}
func (r *PublicRoutes) RegisterPublicRoutes(app *fiber.App) {
	public := app.Group("/public")
	public.Post("", r.PublicSchemaSeed)
	public.Post("/setting_types", r.SettingTypesBulkCreate)
	public.Post("/settings", r.SettingsBulkCreate)
	public.Post("/icons", r.IconsBulkCreate)

}
