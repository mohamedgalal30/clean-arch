package usecase

import (
	"github.com/fly365com/flybase"
	"ticket/entity"
	"ticket/interfaces/ms"
	"ticket/interfaces/repository"
)

type IRadarMs interface {
	LogCreateTicketActivity(ticket *entity.Ticket)
}
type TicketUsecase struct {
	ticketRepo entity.ITicketRepo
	radarMs    IRadarMs
	//Logger     Logger
}

func NewTicketUsecase() *TicketUsecase {
	ticketRepo := repository.NewDbTicketRepo()
	radarMs := ms.NewRadarMs()
	return &TicketUsecase{
		ticketRepo: ticketRepo,
		radarMs:    radarMs,
	}
}

type Logger interface {
	Log(args ...interface{})
}

type Ticket struct {
	Id   string
	Body string
}

func (ucase *TicketUsecase) Add(number int, email string, body string) error {

	ticket := entity.NewTicket(body)
	err := ucase.ticketRepo.Add(ticket)
	if err != nil {
		flybase.App().GetLogger().Error(err.Error())
	}
	ucase.radarMs.LogCreateTicketActivity(ticket)
	//ucase.Logger.Log("Ticket Created")
	return nil
}

func (ucase *TicketUsecase) Search() ([]Ticket, error) {
	tickets := ucase.ticketRepo.FindAll()
	customTickets := getCustomTickets(tickets)
	return customTickets, nil
}

func getCustomTickets(tickets []entity.Ticket) []Ticket {
	customTickets := make([]Ticket, len(tickets))
	for i, ticket := range tickets {
		customTickets[i] = Ticket{ticket.Id, ticket.Body}
	}
	return customTickets
}
