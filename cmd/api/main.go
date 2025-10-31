package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/22Fariz22/crm-estate/config"
	"github.com/22Fariz22/crm-estate/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewPsqlDB(ctx, cfg)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer database.CloseDB(db)

	// тест
	var version string
	if err := db.Get(&version, "SELECT version()"); err != nil {
		log.Printf("Query failed: %v", err)
	} else {
		log.Printf("PostgreSQL: %s", version)
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")
}
