package main

import (
	"libraryRestAPI/internal/config"
	"libraryRestAPI/internal/connection"
	"libraryRestAPI/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()
	customerRepository := repository.NewCustomer(dbConnection)
	app.Get("/developers", developers)
	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
