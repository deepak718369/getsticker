package data

import (
	"fmt"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var connection *gorm.DB
var err error

type DBService struct {
}

func GetConnection() {

	databaseUri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err = gorm.Open("postgres", databaseUri)

	if err != nil {
		os.Exit(1)
	}
	connection.DB().SetMaxIdleConns(2)
	connection.DB().SetMaxOpenConns(30)
	connection.DB().SetConnMaxLifetime(time.Hour * time.Duration(5))
	connection.SingularTable(true)

}

// GetDB : Get an instance of DB to connect to the database connection pool
func (d DBService) GetDB() *gorm.DB {
	return connection
}
