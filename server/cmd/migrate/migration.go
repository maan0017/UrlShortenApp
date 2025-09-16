package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

// this function is solly created to handle migration's only.
func main() {
	// 1. check if arguments are passed or not
	if len(os.Args) < 2 {
		log.Fatalln("please provide a direction 'up' or 'down'")
	}

	// 2. get the aurgument
	arguments := os.Args
	for index, _arg := range arguments {
		fmt.Println(index, ": ", _arg)
	}
	direction := arguments[1]
	fmt.Println("direcion >", direction)

	// 3. open db and defer db.Close()
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalln("Initalization DB Error >", err)
	}
	defer db.Close()

	// 4. create a instance
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalln("DB Instance Error >", err)
	}
	defer instance.Close()

	// 5. get file source
	fileSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal("[ERROR]::File Open Error >", err)
	}

	// 6. create instance
	_migrate, err := migrate.NewWithInstance("file", fileSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal("[ERROR]::Migration Initializing Error >", err)
	}

	switch direction {
	case "up":
		if err := _migrate.Up(); err != migrate.ErrNoChange {
			log.Fatal("[ERROR]::Migration Direction Error >", err)
		}
	case "down":
		if err := _migrate.Down(); err != migrate.ErrNoChange {
			log.Fatal("[ERROR]::Migration Direction Error >", err)
		}
	default:
		log.Fatal("Invalid direction. Use 'up' or 'down'")
	}
	fmt.Println("")
}
