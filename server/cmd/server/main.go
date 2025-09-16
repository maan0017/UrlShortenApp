package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/maan0017/UrlShortner/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1.load config
	config := config.NewConfig()
	config.LoadConfig()

	// 2.load db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error while opening db.", err)
	}
	defer db.Close()

	// 3.load modules and repos

	// 4.load routes

	// 5. start server
	log.Println("server is running at port ", config.Port)
	if err := http.ListenAndServe(config.Port, nil); err != nil {
		log.Fatalln("error while running server: ", err)
	}
}
