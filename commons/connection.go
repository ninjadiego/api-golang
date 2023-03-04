package commons

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Getconnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?chatsset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db.Getconnection()
	defer db.Close()

	log.Println("migrando....")

	db.AutoMigrate(&models.persona{})
}
