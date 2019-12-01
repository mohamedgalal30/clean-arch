package api

import (
	"github.com/fly365com/flybase"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"ticket/infrastructure"
	"ticket/interfaces/api/routes"
	"ticket/interfaces/repository"
)

func Main() {
	app := flybase.App()
	dbHandler := infrastructure.GetTicketDbConnection()
	infrastructure.MigrateTicketPgDB()
	repository.InitDbTicketRepo(dbHandler)
	routes.InitControllers()
	app.SetRoutes(routes.GetRoutes())
	app.StartServer()
}
