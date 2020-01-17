package main

import (
	"flag"
	"log"
	"net/http"
	"simpleGoWeb/process_configs"
)

var assetsPath string

func processFlags() *process_configs.Config {
	cfg := &process_configs.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:8001", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/posgresql " +
		"dbname=simplegoweb sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *process_configs.Config)  {
	log.Printf("Обслуживание отрисовки из %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := process_configs.Run(cfg); err != nil {
		log.Printf("Ошибка в main(): %v", err)
	}
}
