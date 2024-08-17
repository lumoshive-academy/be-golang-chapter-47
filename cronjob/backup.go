package cronjob

import (
	"be-golang-chapter-47/config"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func BackupDatabase(cfg *config.Config) {
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("%s/backup_%s.sql", cfg.BackupDir, timestamp)

	cmd := exec.Command("pg_dump",
		"-h", cfg.DBHost,
		"-p", cfg.DBPort,
		"-U", cfg.DBUser,
		"-d", cfg.DBName,
		"-f", backupFile,
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", cfg.DBPassword))

	err := cmd.Run()
	if err != nil {
		log.Printf("Backup failed: %v", err)
	} else {
		log.Printf("Backup successful: %s", backupFile)
	}
}
