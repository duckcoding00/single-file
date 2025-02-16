package api

import "github.com/duckcoding00/single-file/internal/handler"

func InitServer() {
	handler := handler.NewHandler()
	config := AppConfig{
		handler: handler,
		addr:    ":8080",
	}

	app := NewApp(config)
	app.RegisterRoute()
	app.Run()
}
