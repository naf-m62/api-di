package main

import (
	"api-di/apps/server"
	"api-di/services"
	"github.com/sarulabs/di"
)

func main() {
	defer services.Logger.Sync()

	app := *buildConteiners()
	defer app.Delete()

	server.Start(app)
}

func buildConteiners() *di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		services.Logger.Fatal(err.Error())
	}

	err = builder.Add(services.Services...)
	if err != nil {
		services.Logger.Fatal(err.Error())
	}
	app := builder.Build()
	return &app
}
