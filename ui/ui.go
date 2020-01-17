package ui

import (
	"fmt"
	"net"
	"net/http"
	"simpleGoWeb/model"
	"time"
)


type Config struct {
	// Обслуживание статики
	Assets http.FileSystem
}

const indexHTML = `
<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Просто веб приложение на Go</title>
  </head>
  <body>
    <div id='root'></div>
  </body>
</html>
`

// Не сам обработчик! Возвращает функцию-обработчик
// Для того, что бы *model.Model можно было передать через замыкание
func indexHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, indexHTML)
	})
}

func Start(cfg Config, m *model.Model, listener net.Listener) {
	server := &http.Server{
		ReadTimeout:60 * time.Second,
		WriteTimeout:60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/", indexHandler(m))

	go server.Serve(listener)
}
