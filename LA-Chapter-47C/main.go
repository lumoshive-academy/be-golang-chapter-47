// predefined  schedules
package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Menjalankan tugas setiap jam
	c.AddFunc("@hourly", func() {
		fmt.Println("Tugas dijalankan setiap jam")
	})

	// Menjalankan tugas setiap hari pada jam 00:00
	c.AddFunc("@daily", func() {
		fmt.Println("Tugas dijalankan setiap hari pada jam 00:00")
	})

	// Menjalankan tugas setiap minggu pada hari Minggu pukul 00:00
	c.AddFunc("@weekly", func() {
		fmt.Println("Tugas dijalankan setiap minggu pada hari Minggu pukul 00:00")
	})

	// Menjalankan tugas setiap bulan pada hari pertama pukul 00:00
	c.AddFunc("@monthly", func() {
		fmt.Println("Tugas dijalankan setiap bulan pada hari pertama pukul 00:00")
	})

	// Menjalankan tugas setiap tahun pada tanggal 1 Januari pukul 00:00
	c.AddFunc("@yearly", func() {
		fmt.Println("Tugas dijalankan setiap tahun pada tanggal 1 Januari pukul 00:00")
	})

	// Memulai cron scheduler
	c.Start()

	// Agar program tidak langsung selesai
	select {}
}

// // interval
// package main

// import (
//     "fmt"
//     "github.com/robfig/cron/v3"
//     "time"
// )

// func main() {
//     c := cron.New()

//     // Menjalankan tugas setiap 5 menit
//     c.AddFunc("@every 5m", func() {
//         fmt.Println("Tugas dijalankan setiap 5 menit")
//     })

//     // Menjalankan tugas setiap 2 jam
//     c.AddFunc("@every 2h", func() {
//         fmt.Println("Tugas dijalankan setiap 2 jam")
//     })

//     // Menjalankan tugas setiap hari pada pukul 1 pagi dan setiap 30 menit
//     c.AddFunc("@every 30m", func() {
//         fmt.Println("Tugas dijalankan setiap 30 menit")
//     })

//     // Menjalankan tugas setiap 10 detik
//     c.AddFunc("@every 10s", func() {
//         fmt.Println("Tugas dijalankan setiap 10 detik")
//     })

//     // Memulai cron scheduler
//     c.Start()

//     // Agar program tidak langsung selesai
//     select {}
// }
