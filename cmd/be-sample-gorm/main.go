package main

import (
	"fmt"
	rest "github.com/HiBang15/sample-gorm.git/api/rest/be-sample-gorm"
	"github.com/HiBang15/sample-gorm.git/api/rest/be-sample-gorm/router/public"
	"github.com/HiBang15/sample-gorm.git/internal/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseInfo := &database.DatabaseInfo{
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	//connect database
	err = database.ConnectToDB(databaseInfo)
	if err != nil {
		log.Fatalf("Connect database fail with error: %v", err.Error())
	}
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	//run Rest API
	fmt.Println("Start run REST API OF TWC ADMIN API .....")
	settingTwcAminApi := rest.SettingRestApi{
		Environment: os.Getenv("ENVIRONMENT"),
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
	}

	rest.Load(settingTwcAminApi, public.SetRouter)
}
