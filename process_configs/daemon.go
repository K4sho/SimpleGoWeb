package process_configs

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"simpleGoWeb/db"
	"simpleGoWeb/model"
	"simpleGoWeb/ui"
)

type Config struct {
	ListenSpec string

	Db db.Config
	UI ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Запуск, HTTP на: %s\n", cfg.ListenSpec)

	db, err := db.InitDB(cfg.Db)
	if err != nil {
		log.Printf("Ошибка в инициализации БД: %v\n", err)
		return err
	}

	m := model.New(db)

	// Прослушивание tcp порта
	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Не удалось создать слушателя: %v\n", err)
		return err
	}

	ui.Start(cfg.UI, m, l)

	waitForSignal()

	return nil
}

func waitForSignal()  {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Получен сигнал: %v, exiting.", s)
}