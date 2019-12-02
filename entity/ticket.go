package entity

import (
	"github.com/fly365com/flybase"
)

type ITicketDbRepo interface {
	Add(ticket *Ticket) error
	FindAll() []Ticket
}

type Ticket struct {
	flybase.PgDbModel
	Number int64  `gorm:"column:number" json:"number" `
	Email  string `gorm:"type:varchar(255);column:email" json:"email"`
	Body   string `gorm:"type:text; column:body" json:"body"`
}

func NewTicket(body string) *Ticket {
	var ticket Ticket
	ticket.Body = body
	return &ticket
}
