package db

import (
	"github.com/jinzhu/gorm"
	"ticket/entity"
)

type DbHandler *gorm.DB

type DbRepo struct {
	DbHandler *gorm.DB
}

type TicketDbRepo DbRepo

var dbHandler *gorm.DB

func InitDbTicketRepo(dbHndlr *gorm.DB) {
	dbHandler = dbHndlr
}

func NewDbTicketRepo() *TicketDbRepo {
	dbTicketRepo := new(TicketDbRepo)
	dbTicketRepo.DbHandler = dbHandler
	return dbTicketRepo
}

func (repo *TicketDbRepo) Add(ticket *entity.Ticket) error {
	err := repo.DbHandler.Create(ticket)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *TicketDbRepo) FindAll() []entity.Ticket {
	return []entity.Ticket{}
}
