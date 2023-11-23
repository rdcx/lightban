package main

import (
	"lightban/api/db"
	"lightban/api/handler"
	"lightban/api/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	db := db.NewDB(dsn)

	h := handler.NewHandler(db)

	r := router.SetUp(h)
	r.Run(":8080")
}
