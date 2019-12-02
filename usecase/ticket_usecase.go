package usecase

import (
	"github.com/fly365com/flybase"
	"ticket/entity"
	"ticket/repository/db"
	"ticket/repository/ms"
)

type IRadarMsRepo interface {
	LogCreateTicketActivity(ticket *entity.Ticket)
}
type TicketUsecase struct {
	ticketRepo entity.ITicketDbRepo
	radarMs    IRadarMsRepo
}

func NewTicketUsecase() *TicketUsecase {
	ticketRepo := db.NewDbTicketRepo()
	radarMs := ms.NewRadarMs()
	return &TicketUsecase{
		ticketRepo: ticketRepo,
		radarMs:    radarMs,
	}
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
