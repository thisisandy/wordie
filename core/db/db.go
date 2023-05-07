package db

import (
	"log"
	"os"
	"strings"
	"sync"
	"wordie/core/db/wordDB"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wordie/core/db/userDB"

	"github.com/joho/godotenv"
)

func loadEnv() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dict := make(map[string]string)
	for _, v := range os.Environ() {
		pair := strings.Split(v, "=")
		dict[pair[0]] = pair[1]
	}
	return dict

}

var database *gorm.DB
var once sync.Once

func Instance() *gorm.DB {
	once.Do(func() {
		env := loadEnv()
		dsn := env["DB_USERNAME"] + ":" + env["DB_PASSWORD"] + "@tcp(" + env["DB_HOST"] + ":" + env["DB_PORT"] + ")/" + env["DB_NAME"] + "?charset=utf8mb4&parseTime=True&loc=Local"
		_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		database = _db
		userDB.CreateUserTable(database)
		wordDB.CreateWordTable(database)
	})
	return database
}
