package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Db *gorm.DB
}

func (pg *PostgreSQL) Connect() {
	dsn := "host=localhost user=root password=secret dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	pg.Db = db.Debug()

}
