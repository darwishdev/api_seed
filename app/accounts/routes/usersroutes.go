package routes

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/meloneg/mln_data_pool/common/file"
	"github.com/rs/zerolog/log"
)

type RequestBody struct {
	Test bool `json:"test"`
}

var tempPath = "static/temp/"

func (r *AccountRoutes) PermissionsBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	permissions, err := r.service.PermissionsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(permissions)
}

func (r *AccountRoutes) RolesBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	roles, err := r.service.RolesBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(roles)
}

func (r *AccountRoutes) AuthUsersBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	roles, err := r.service.AuthUsersBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(roles)
}
func (r *AccountRoutes) RolesBulkCreateFile(c *fiber.Ctx) error {
	// Parse the form data to get the uploaded file
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var requestFile *multipart.FileHeader
	for _, files := range form.File {
		requestFile = files[0]
	}
	fileName := filepath.Join(tempPath, requestFile.Filename)

	if err := handleFileUpload(requestFile, fileName); err != nil {
		return err
	}
	defer handleFileRemove(requestFile, fileName)

	uploadedFile, err := file.LoadFile(fileName)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load users file")
	}

	rows, err := uploadedFile.GetRows("roles")
	if err != nil {
		return err
	}

	count, err := r.service.RolesBulkCreateFile(c.Context(), rows)

	log.Debug().Interface("count", count).Msg("created")
	if err != nil {
		return err
	}
	return c.JSON(count)
}

func handleFileUpload(file *multipart.FileHeader, name string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a destination file
	dst, err := os.Create(name)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func handleFileRemove(file *multipart.FileHeader, name string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Delete the source file
	if err = os.Remove(name); err != nil {
		return err
	}

	return nil
}

func (r *AccountRoutes) UsersBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	users, err := r.service.UsersBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
func (r *AccountRoutes) CustomersBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	customers, err := r.service.CustomersBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(customers)
}

func (r *AccountRoutes) OwnersBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	users, err := r.service.OwnersBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (r *AccountRoutes) NavigationBarsBulkCreate(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	navigations, err := r.service.NavigationBarsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(navigations)
}

func (r *AccountRoutes) UsersSchemaSeed(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		requestBody.Test = false
	}
	permissions, err := r.service.PermissionsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	roles, err := r.service.RolesBulkCreate(c.Context())
	if err != nil {
		return err
	}
	users, err := r.service.UsersBulkCreate(c.Context())
	if err != nil {
		return err
	}
	navigations, err := r.service.NavigationBarsBulkCreate(c.Context())
	if err != nil {
		return err
	}
	response := map[string]any{
		"permissions": permissions,
		"roles":       roles,
		"users":       users,
		"navigations": navigations,
	}
	return c.JSON(response)
}
