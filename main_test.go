package datatypes_test

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	log.Println("testing mysql...")
	if dbDSN == "" {
		dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=UTC"
	}
	db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func SupportedDriver(dialectors ...string) bool {
	for _, dialect := range dialectors {
		if DB.Dialector.Name() == dialect {
			return true
		}
	}
	return false
}
