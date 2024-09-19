package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meloneg/mln_data_pool/app/storage/service"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supaclient"
)

type StorageRoutesInterface interface {
	RegisterStoragesRoutes(app *fiber.App)
}

type StorageRoutes struct {
	service service.StorageServiceInterface
	store   db.Store
}

func NewStorageRoutes(store db.Store, supa supaclient.SupabaseServiceInterface) StorageRoutesInterface {
	userService := service.NewStorageService(supa)
	return &StorageRoutes{
		service: userService,
		store:   store,
	}
}
func (r *StorageRoutes) RegisterStoragesRoutes(app *fiber.App) {
	storage := app.Group("/storage")
	storage.Post("/", r.SeedStorage)
}
