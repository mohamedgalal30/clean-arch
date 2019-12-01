package infrastructure

import (
	"errors"
	"fmt"
	"github.com/fly365com/flybase"
	"github.com/jinzhu/gorm"
	"ticket/entity"
)

type PostgresHandler struct {
	DB *gorm.DB
}

func GetTicketDbConnection() *gorm.DB {
	const testDbName = "test"
	db := flybase.GetPgDbConnectionByName(testDbName)
	if db.Error != nil {
		flybase.MyApp.GetLogger().Error(db.Error)
	}
	return db
}
func MigrateTicketPgDB() {
	conn := GetTicketDbConnection()

	if err := conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" ").Error; err != nil {
		fmt.Println("uuid-ossp creating extention failed")
		panic(fmt.Sprintf("uuid-ossp creating extention failed %+v", err))
	}

	db := conn.AutoMigrate(&entity.Ticket{})
	if db.Error != nil {
		flybase.MyApp.GetLogger().Error(errors.New("Error while creating Ticket table"))
		flybase.MyApp.GetLogger().Error(db.Error)
	}
}
