package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/daniilcdev/back-office-ws/config"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.LoadEnv()
	db, err := sql.Open(cfg.DbAdapter, cfg.DbConnection)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("db connected")

	r := chi.NewRouter()
	r.Get("/", http.FileServer(http.Dir("./static")).ServeHTTP)

	log.Println("starting server...")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Println(err)
	}

	waitForExit()
}

func waitForExit() {
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)
	<-interupt
}
