package routes

import (
	"github.com/fly365com/flybase"
	"ticket/delivery/api/controller"
)

var TicketCtrl controller.TicketController

func InitControllers() {
	TicketCtrl = *controller.NewTicketController()
}

func GetRoutes() []flybase.Route {
	var appRoutes []flybase.Route
	appRoutes = append(appRoutes, TicketRoutes...)
	return appRoutes
}
