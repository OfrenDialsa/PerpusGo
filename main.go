package main

import (
	"PerpusGo/internal/api"
	"PerpusGo/internal/config"
	"PerpusGo/internal/connection"
	"PerpusGo/internal/repository"
	"PerpusGo/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
