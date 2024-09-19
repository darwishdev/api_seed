package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meloneg/mln_data_pool/app/accounts/factory"
	"github.com/meloneg/mln_data_pool/app/accounts/repo"
	"github.com/meloneg/mln_data_pool/app/accounts/service"
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/config"
	"github.com/meloneg/mln_data_pool/supaclient"
	"github.com/xuri/excelize/v2"
)

type AccountRoutesInterface interface {
	RegisterUsersRoutes(app *fiber.App)
}

type AccountRoutes struct {
	service service.AccountServiceInterface
}

func NewAccountRoutes(store db.Store, data *excelize.File, config config.Config, supa supaclient.SupabaseServiceInterface) AccountRoutesInterface {
	accountsFactory := factory.NewAccountsFactory(data, supa)
	accountRepo := repo.NewAccountRepo(store)
	accountService := service.NewAccountService(data, accountsFactory, accountRepo, supa)
	return &AccountRoutes{
		service: accountService,
	}
}
func (r *AccountRoutes) RegisterUsersRoutes(app *fiber.App) {
	app.Static("/templates", "./data/templates/")
	accounts := app.Group("/accounts")
	accounts.Post("/", r.UsersSchemaSeed)
	accounts.Post("/users", r.UsersBulkCreate)
	accounts.Post("/customers", r.CustomersBulkCreate)
	accounts.Post("/owners", r.OwnersBulkCreate)
	accounts.Post("/authusers", r.AuthUsersBulkCreate)
	accounts.Post("/navigations", r.NavigationBarsBulkCreate)
	accounts.Post("/permissions", r.PermissionsBulkCreate)
	accounts.Post("/roles", r.RolesBulkCreate)
	accounts.Post("/roles/file", r.RolesBulkCreateFile)
}
