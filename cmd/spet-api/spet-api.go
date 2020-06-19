package main

import (
	"SPET/internal/spet-api/api"
	"SPET/internal/spet-api/settings"
	"SPET/pkg/lig"
	"log"
)

func main() {
	// Инициализация системы логирования
	err := lig.Create("logs.txt")
	if err != nil {
		// В случаи не отсутствия доступа к файлу хранения логов закроем приложение
		log.Fatal("[CRASH] Logger system isn't init ->", err)
	}
	defer lig.Quit()

	// Настройки по умолчанию
	defaultConfig := &settings.Config{
		PathToLogger: "logs.txt",
		Address:      "3000",
		Database: settings.Database{
			Name: "postgres",
			URL:  "host=localhost password=1234 dbname=spet sslmode=disable",
		},
	}

	// Инициализация настроек
	config := &settings.Config{}
	config, err = settings.NewConfig(defaultConfig, "./config/config.json")
	if err != nil {
		// В случаи не отсутствия доступа к файлу хранения логов закроем приложение
		lig.Crash("Config file isn't created", err)
	}

	// Инициализация сервера
	s := api.New(config)
	if err := s.Start(); err != nil {
		lig.Crash("Server is not started", err)
	}
}
