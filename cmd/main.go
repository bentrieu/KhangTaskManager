package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Task struct {
	TaskId    string
	Descript  string
	Progress  string
	CreatedAt string
}

func getData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get data
		rawData, err := db.Query("SELECT * FROM tasks")
		if err != nil {
			log.Fatal("error getting data", err)
		}
		defer rawData.Close()

		data := make([]*Task, 0)
		for rawData.Next() {
			task := &Task{}
			err := rawData.Scan(&task.TaskId, &task.Descript, &task.Progress, &task.CreatedAt)
			if err != nil {
				log.Fatal("error scannign data", err)
			}
			data = append(data, task)
		}
		enc := json.NewEncoder(w)
		enc.Encode(data)
		if err != nil {
			log.Fatal("error encoding data", err)
		}
	}
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
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("database migration error", err)
	}
	//setup http server
	//start
	http.HandleFunc("/data", getData(db))

	fmt.Println("running!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
