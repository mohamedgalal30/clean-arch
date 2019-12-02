package routes

import (
	"github.com/fly365com/flybase"
	"net/http"
)

var TicketRoutes = []flybase.Route{
	{http.MethodGet,
		"/tickets",
		TicketCtrl.List,
	},
	{http.MethodPost,
		"/tickets",
		TicketCtrl.Create,
	},
}
