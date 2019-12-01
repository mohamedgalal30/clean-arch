package repository

import (
	"github.com/jinzhu/gorm"
	"ticket/entity"
)

type DbHandler *gorm.DB

type DbRepo struct {
	DbHandler *gorm.DB
}

type DbTicketRepo DbRepo

var dbHandler *gorm.DB

func InitDbTicketRepo(dbHndlr *gorm.DB) {
	dbHandler = dbHndlr
}

func NewDbTicketRepo() *DbTicketRepo {
	dbTicketRepo := new(DbTicketRepo)
	dbTicketRepo.DbHandler = dbHandler
	return dbTicketRepo
}

func (repo *DbTicketRepo) Add(ticket *entity.Ticket) error {
	err := repo.DbHandler.Create(ticket)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *DbTicketRepo) FindAll() []entity.Ticket {
	return []entity.Ticket{}
}
