package controller

import (
	"encoding/json"
	"github.com/fly365com/flybase"
	"io/ioutil"
	"net/http"
	"ticket/entity"
	"ticket/delivery/api/errors"
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
	//request := self.NewListRequest(r, reflect.ValueOf(&ticket).Elem())
	//dataSource := flybase.GetDataSource(r)
	//defaultDataSource := flybase.GetDefaultDataSource()
	//flybase.App().GetLogger().Infof("default data source %s", defaultDataSource)
	//flybase.App().GetLogger().Infof("data source %s", dataSource)
	//query := r.URL.Query()
	//prioritiesIds := query["priority"]
	//categoriesIds := query["category"]
	//statusIds := query["status"]
	//tagsIds := query["tag"]
	//tickets, total, resultDataSource := ticket.FindAllFactory(
	//	request, dataSource, prioritiesIds, statusIds, categoriesIds, tagsIds)
	//response := request.GetResponse()
	//response.Data = tickets
	//response.Total = total
	//response.Update()
	//w = flybase.SetDataSource(w, resultDataSource)
	ctrl.Json(w, ticket, http.StatusOK)
}

func (ctrl *TicketController) Create(w http.ResponseWriter, req *http.Request) {
	//rawStrings, err := TicketUnmarshalJSON(req)
	err := ctrl.TicketUcase.Add(1, "mohamed.galal@fly365.com", "Body Body")
	if err != nil {
		flybase.App().GetLogger().Error(err)
		ctrl.JsonError(w, errors.BodyTypeError, http.StatusBadRequest)
		return
	}

	ctrl.Json(w, "added", http.StatusCreated)
}

func TicketUnmarshalJSON(req *http.Request) (map[string]interface{}, error) {
	var rawStrings map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &rawStrings)
	if err != nil {
		return nil, err
	}
	return rawStrings, nil
}
