// cron expression charater
package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Menjadwalkan tugas sesuai dengan ekspresi CRON yang telah dijelaskan
	c.AddFunc("15 30 10 5 1 1", func() {
		fmt.Println("Task dijalankan pada 5 Januari, pukul 10:30:15 AM, hari Senin")
	})

	c.AddFunc("45 15 14 12 2 2", func() {
		fmt.Println("Task dijalankan pada 12 Februari, pukul 02:15:45 PM, hari Selasa")
	})

	c.AddFunc("0 0 8 25 12 3", func() {
		fmt.Println("Task dijalankan pada 25 Desember, pukul 08:00:00 AM, hari Rabu")
	})

	c.AddFunc("30 45 18 30 6 4", func() {
		fmt.Println("Task dijalankan pada 30 Juni, pukul 06:45:30 PM, hari Kamis")
	})

	c.Start()

	// Agar program tidak langsung selesai
	select {}
}

// //special character
// package main

// import (
//     "fmt"

//     "github.com/robfig/cron/v3"
// )

// func main() {
//     c := cron.New()

//     // Asterisk (*): Menjalankan tugas setiap menit
//     c.AddFunc("* * * * *", func() {
//         fmt.Println("Tugas dijalankan setiap menit")
//     })

//     // Comma (,): Menjalankan tugas pada jam 8:00 AM dan 6:00 PM setiap hari
//     c.AddFunc("0 8,18 * * *", func() {
//         fmt.Println("Tugas dijalankan pada jam 8:00 AM dan 6:00 PM setiap hari")
//     })

//     // Dash (-): Menjalankan tugas setiap jam dari 9:00 AM hingga 5:00 PM setiap hari
//     c.AddFunc("0 9-17 * * *", func() {
//         fmt.Println("Tugas dijalankan setiap jam dari 9:00 AM hingga 5:00 PM setiap hari")
//     })

//     // Slash (/): Menjalankan tugas setiap 5 menit
//     c.AddFunc("*/5 * * * *", func() {
//         fmt.Println("Tugas dijalankan setiap 5 menit")
//     })

//     // Question Mark (?): Menjalankan tugas pada jam 12:00 PM setiap hari
//     c.AddFunc("0 0 12 * * ?", func() {
//         fmt.Println("Tugas dijalankan pada jam 12:00 PM setiap hari")
//     })

//     // Memulai cron scheduler
//     c.Start()

//     // Agar program tidak langsung selesai
//     select {}
// }
