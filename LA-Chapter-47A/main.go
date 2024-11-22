package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Menjadwalkan tugas untuk dijalankan setiap menit
	c.AddFunc("* * * * *", func() {
		fmt.Println("Task dijalankan setiap menit")
	})

	// Memulai scheduler
	c.Start()

	// Agar program tidak langsung selesai
	select {}
}
