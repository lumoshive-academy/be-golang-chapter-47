// timezone
package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithLocation(time.UTC)) // Set zona waktu default ke UTC

	// Menjalankan tugas setiap hari pada pukul 12:00 PM UTC
	c.AddFunc("0 12 * * *", func() {
		fmt.Println("Tugas dijalankan setiap hari pada pukul 12:00 PM UTC")
	})

	// Menjalankan tugas dengan zona waktu lokal yang berbeda
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	c = cron.New(cron.WithLocation(loc))

	// Menjalankan tugas setiap hari pada pukul 12:00 PM WIB
	c.AddFunc("0 12 * * *", func() {
		fmt.Println("Tugas dijalankan setiap hari pada pukul 12:00 PM WIB")
	})

	// Memulai cron scheduler
	c.Start()

	// Agar program tidak langsung selesai
	select {}
}

// // logging
// package main

// import (
//     "fmt"
//     "log"
//     "os"

//     "github.com/robfig/cron/v3"
// )

// // Kustom logger yang mengimplementasikan cron.Logger
// type customLogger struct {
//     logger *log.Logger
// }

// // Info logs an informational message with optional arguments
// func (l *customLogger) Info(msg string, args ...interface{}) {
//     l.logger.Printf("[INFO] "+msg, args...)
// }

// // Error logs an error message with an error and optional arguments
// func (l *customLogger) Error(err error, msg string, args ...interface{}) {
//     l.logger.Printf("[ERROR] "+msg+": %v", append(args, err)...)
// }

// func main() {
//     // Membuat logger kustom
//     file, err := os.OpenFile("cron.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//     if err != nil {
//         fmt.Println("Error creating log file:", err)
//         return
//     }
//     defer file.Close()

//     logger := &customLogger{
//         logger: log.New(file, "", log.LstdFlags),
//     }

//     // Membuat scheduler cron dengan logger kustom
//     c := cron.New(cron.WithLogger(logger))

//     // Menambahkan tugas
//     c.AddFunc("*/1 * * * *", func() {
//         fmt.Println("Tugas dijalankan setiap menit")
//     })

//     // Memulai cron scheduler
//     c.Start()

//     // Agar program tidak langsung selesai
//     select {}
// }
