package oldNotifier

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
		if err != nil {
			panic(err)
		}
		instance = db
	})
	return instance
}
