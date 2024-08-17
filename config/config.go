package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	BackupDir    string
	CronSchedule string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		DBHost:       viper.GetString("database.host"),
		DBPort:       viper.GetString("database.port"),
		DBUser:       viper.GetString("database.user"),
		DBPassword:   viper.GetString("database.password"),
		DBName:       viper.GetString("database.name"),
		BackupDir:    viper.GetString("backup.dir"),
		CronSchedule: viper.GetString("cron.schedule"),
	}

	return config, nil
}
