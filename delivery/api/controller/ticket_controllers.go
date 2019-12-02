package controller

import (
	"github.com/fly365com/flybase"
	"net/http"
	"ticket/delivery/api/errors"
	"ticket/entity"
	"ticket/usecase"
)

type ITicketUsecase interface {
	Search() ([]usecase.Ticket, error)
	Add(number int, email string, body string) error
}

type TicketController struct {
	flybase.Controller
	TicketUcase ITicketUsecase
}

func NewTicketController() *TicketController {
	ticketUcase := usecase.NewTicketUsecase()
	return &TicketController{
		TicketUcase: ticketUcase,
	}
}

func (ctrl *TicketController) List(w http.ResponseWriter, r *http.Request) {

	var ticket entity.Ticket
	ctrl.Json(w, ticket, http.StatusOK)
}

func (ctrl *TicketController) Create(w http.ResponseWriter, req *http.Request) {
	err := ctrl.TicketUcase.Add(1, "mohamed.galal@fly365.com", "Body Body")
	if err != nil {
		flybase.App().GetLogger().Error(err)
		ctrl.JsonError(w, errors.BodyTypeError, http.StatusBadRequest)
		return
	}

	ctrl.Json(w, "added", http.StatusCreated)
}
