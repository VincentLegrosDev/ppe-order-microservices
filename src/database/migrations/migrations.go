package main

import (
	"embed"
	"github.com/pressly/goose/v3"
	"os"
	"log"
	"ppe4peeps.com/services/database/internal"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDatabase();   

	goose.SetBaseFS(embedMigrations) 

	if err := goose.SetDialect(os.Getenv("DB_DRIVER")); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
