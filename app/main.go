package main

import (
	"fmt"
	"os"
	"time"

	config "github.com/getsticker/config"
	domain "github.com/getsticker/domain"
	_stickerHttpDelivery "github.com/getsticker/sticker/delivery/http"
	_stickerRepo "github.com/getsticker/sticker/repository/postgres"
	_stickerucase "github.com/getsticker/sticker/usecase"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	logg "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	config.InitConfig()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	connection := GetConnection()

	ar := _stickerRepo.NewpostgresStickerRepository(connection)
	au := _stickerucase.NewStickerUsecase(ar)
	_stickerHttpDelivery.NewstickerHandler(e, au)

	logg.Fatal(e.Start(viper.GetString(`server.address`)))
}

// GetConnection make podtgress connection
func GetConnection() *gorm.DB {

	dbHost := viper.GetString(`database.DB_Host`)
	dbPort := viper.GetString(`database.DB_Port`)
	dbUser := viper.GetString(`database.DB_User`)
	dbPass := viper.GetString(`database.DB_Password`)
	dbName := viper.GetString(`database.DB_Name`)

	databaseURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	connection, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		os.Exit(1)
	}

	connection.DB().SetMaxIdleConns(2)
	connection.DB().SetMaxOpenConns(30)
	connection.DB().SetConnMaxLifetime(time.Hour * time.Duration(5))
	connection.SingularTable(true)

	migrationErr := connection.AutoMigrate(&domain.Sticker{}).Error
	if migrationErr != nil {
		os.Exit(1)
	}
	return connection
}
