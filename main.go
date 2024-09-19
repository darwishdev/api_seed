package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	accountRoutes "github.com/meloneg/mln_data_pool/app/accounts/routes"

	publicRoutes "github.com/meloneg/mln_data_pool/app/public/routes"
	storegeRoutes "github.com/meloneg/mln_data_pool/app/storage/routes"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/common/file"
	"github.com/meloneg/mln_data_pool/config"
	"github.com/meloneg/mln_data_pool/supaclient"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
)

func loafFiles(fileNames map[string]string) (map[string]*excelize.File, error) {
	files := make(map[string]*excelize.File, len(fileNames))
	for key, value := range fileNames {
		file, err := file.LoadFile(value)
		if err != nil {
			return nil, err
		}
		files[key] = file
	}
	return files, nil
}

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) < 2 {
		panic("please provide env path as well as excel folder path")
	}

	var (
		app             = fiber.New()
		ctx             = context.Background()
		envFilePath     = os.Args[1]
		excelFolderPath = os.Args[2]
		fileNames       = map[string]string{
			"accounts": fmt.Sprintf("%s/accounts.xlsx", excelFolderPath),
			"public":   fmt.Sprintf("%s/public.xlsx", excelFolderPath),
			"places":   fmt.Sprintf("%s/places.xlsx", excelFolderPath),
			"products": fmt.Sprintf("%s/products.xlsx", excelFolderPath),
			// "stock":    fmt.Sprintf("%s/stock.xlsx", excelFolderPath),
		}
	)

	config, err := config.LoadConfig(envFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the config")
	}
	store, err := db.InitDB(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	files, err := loafFiles(fileNames)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load users file")
	}

	supa, err := supaclient.NewSupabaseService(config.SupaUrl, config.SupaKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to supabase api")
	}
	// repos

	accountRouter := accountRoutes.NewAccountRoutes(store, files["accounts"], config, supa)
	storageRouter := storegeRoutes.NewStorageRoutes(store, supa)
	publicRouter := publicRoutes.NewPublicRoutes(store, files["public"])

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Add the allowed origins here
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	accountRouter.RegisterUsersRoutes(app)
	storageRouter.RegisterStoragesRoutes(app)
	publicRouter.RegisterPublicRoutes(app)
	err = app.Listen(config.HTTPServerAddress)
	if err != nil {
		panic(err)
	}

}
