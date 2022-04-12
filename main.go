package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hweihwang/go-blogs/component"
	"github.com/hweihwang/go-blogs/modules/blog/blogtransport/fiberblog"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CON_STR")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	appCtx := component.NewAppContext(db)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	blogs := v1.Group("/blogs")
	{
		blogs.Post("/", fiberblog.CreateBLog(appCtx))
		blogs.Get("/:id", fiberblog.GetBlog(appCtx))
		blogs.Post("/list", fiberblog.ListBlog(appCtx))
	}

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
