package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	filePath := flag.String("file", "", "path to the sqlite3 database file")
	apiKey := flag.String("api_key", "", "API key to set for admin user")
	flag.Parse()

	if *filePath == "" || *apiKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", *filePath)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	defer db.Close()

	query := `
	UPDATE users 
	SET api_key = ?,
	password_change_required = false
	WHERE id = 1`

	result, err := db.Exec(query, *apiKey)

	if err != nil {
		log.Fatalf("Error executing query: %s", err)

	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Fatalf("No rows affected")
	}
	fmt.Println("Setup complete")

}
