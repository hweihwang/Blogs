package main

import (
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CON_STR")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&blogmodel.Blog{})

	if err != nil {
		log.Fatal(err)
	}
}
