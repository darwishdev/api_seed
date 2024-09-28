package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *StorageRoutes) SeedStorage(c *fiber.Ctx) error {

	err := r.service.BucketCreateImages(c.Context())
	if err != nil {
	}
	err = r.service.UploadInitialImages(c.Context())
	if err != nil {
		return err
	}

	return c.JSON("uploaded")
}
