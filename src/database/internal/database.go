package database


import (
	"fmt"
	"os"
	"log"
	"database/sql"
	 _ "github.com/lib/pq"
)


type Database struct {
	url  string
	DB *sql.DB 
}

func InitDatabase() *sql.DB {  
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", 
			os.Getenv("DB_DRIVER"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

	db,err := sql.Open(os.Getenv("DB_DRIVER"), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
} 


