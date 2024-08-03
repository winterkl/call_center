package main

import (
	"contact_center/internal/app"
	"contact_center/internal/config"
	"log"
)

func main() {
	// Чтение конфига в переменные окружения
	config, err := config.Init()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфигурационных файлов: %s", err.Error())
	}
	// Запуск приложения
	app.NewApp(config).Run()
}
