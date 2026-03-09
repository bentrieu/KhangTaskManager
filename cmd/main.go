package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WRONG"))
}

func main() {
	//config
	// postgres.
	//connect to db
	dsn := "postgres://postgres:example@db:5432/taskmanager?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// ping to check db
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("database connected")

	//init migration
	m, err := migrate.New("file://./changelog", dsn)
	if err != nil {
		log.Fatal("error init database migration", err)
	}

	// run migration
	err = m.Up()
	if err != nil {
		log.Fatal("database migration error", err)
	}
	//setup http server
	//start
	http.HandleFunc("/hello", addHandler)

	fmt.Println("running!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
