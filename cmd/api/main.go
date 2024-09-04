package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
	server *http.ServeMux
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv(""), "postgresSQL")

	flag.Parse()

	fmt.Println(cfg.db.dsn)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Printf("database connection pool established")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: router,
	}

	app := &application{
		config: cfg,
		logger: logger,
		server: server,
	}

}
