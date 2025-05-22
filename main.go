package main

import (
	"PerpusGo/dto"
	"PerpusGo/internal/api"
	"PerpusGo/internal/config"
	"PerpusGo/internal/connection"
	"PerpusGo/internal/repository"
	"PerpusGo/internal/service"
	"net/http"

	jwtMidd "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	jwtMidd := jwtMidd.New(jwtMidd.Config{
		SigningKey: jwtMidd.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("unauthorized, please login"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)
	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository := repository.NewBookStock(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	bookService := service.NewBook(bookRepository, bookStockRepository)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, jwtMidd)
	api.NewBookStock(app, bookStockService, jwtMidd)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
