package api

func InitServer() {
	config := AppConfig{
		addr: ":8080",
	}

	app := NewApp(config)
	app.RegisterRoute()
	app.Run()
}
