package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := os.Getenv("DB_URL")
	dsn := "host=mahmud.db.elephantsql.com user=rtauscwh password=ubPA93lEIaOhwkmLvkfP5DXIZDJ77qpg dbname=rtauscwh port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "rtauscwh:ubPA93lEIaOhwkmLvkfP5DXIZDJ77qpg@mahmud.db.elephantsql.com/rtauscwh"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
