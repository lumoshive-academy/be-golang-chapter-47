package cmd

import (
	"be-golang-chapter-47-implem/service"
	"fmt"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	// Membuat instance cron
	c := cron.New()
	// Menambahkan job untuk membuat laporan setiap hari pukul 9 pagi
	_, err := c.AddFunc("*/2 * * * *", func() {
		service.ReportOrderExcel()
	})

	if err != nil {
		fmt.Println("Error menambahkan cron job:", err)
		return
	}
	// Menjalankan cron
	c.Start()
	// Menunggu agar main tidak langsung selesai
	fmt.Println("Cron job berjalan. Tekan CTRL+C untuk keluar.")
	select {}
}
