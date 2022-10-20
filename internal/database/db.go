package database

import (
	"fmt"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"

	// walletEntities "git02.smartosc.com/smart-wallet/be-smartcontract-wallet/internal/module/wallet/entities"
	"log"
	"sync"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type DatabaseInfo struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSLMode  string `json:"ssl_mode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func Migrate(database *gorm.DB) (err error) {
	err = database.AutoMigrate(
		&entities.User{},
	)
	return err
}

var Connection *gorm.DB
var once sync.Once

func ConnectToDB(database *DatabaseInfo) (err error) {
	once.Do(func() {
		dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			database.Username, database.Password, database.Host, database.Port, database.Name, database.SSLMode)

		//dbURL := "postgres://root:secret@localhost:5432/golang"

		Connection, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
		postgreDB, err := Connection.DB()
		postgreDB.SetMaxIdleConns(20)
		postgreDB.SetMaxOpenConns(200)

		if err != nil {
			log.Printf("Connect database fail with error: %v", err.Error())
		}

		err = Migrate(Connection)
		if err != nil {
			log.Printf("\nMigrate database fail with err: %v", err.Error())
		}
	})
	if err != nil {
		return err
	}

	return nil
}
