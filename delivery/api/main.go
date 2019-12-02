package api

import (
	"github.com/fly365com/flybase"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"ticket/delivery/api/routes"
	"ticket/infrastructure"
	"ticket/repository/db"
)

func Main() {
	app := flybase.App()
	dbHandler := infrastructure.GetTicketDbConnection()
	infrastructure.MigrateTicketPgDB()
	db.InitDbTicketRepo(dbHandler)
	routes.InitControllers()
	app.SetRoutes(routes.GetRoutes())
	app.StartServer()
}
