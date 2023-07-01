package main

import (
	"crm/internal/db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run() error {
	fmt.Println("Running our application")

	dbx, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	fmt.Print("sdfsdf, %w", dbx.Client)

	db.MigrateDatabase(dbx)

	fmt.Println("database migrated")

	return err

}

func main() {
	fmt.Println("Starting the CRM Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)

}
