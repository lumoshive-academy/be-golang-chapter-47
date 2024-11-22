package main

import (
	"be-golang-chapter-47/implem-cron/config"
	"be-golang-chapter-47/implem-cron/cronjob"
	"be-golang-chapter-47/implem-cron/handler"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up cron job
	c := cron.New()
	_, err = c.AddFunc(cfg.CronSchedule, func() { cronjob.BackupDatabase(cfg) })
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}
	c.Start()
	defer c.Stop()

	// Set up web server
	http.HandleFunc("/health", handler.HealthHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
