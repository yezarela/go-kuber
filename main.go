package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"kuba/config"
	"kuba/infra/database"
	_productHandler "kuba/product/handler/http"
	_productRepo "kuba/product/repository"
	_productUsecase "kuba/product/usecase"
)

func init() {
	err := config.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Data.DBHost,
		config.Data.DBPort,
		config.Data.DBUser,
		config.Data.DBName,
		config.Data.DBPassword,
	)

	db, err := database.NewPostgresConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	productRepo := _productRepo.NewRepository(db)
	productUsecase := _productUsecase.NewUsecase(productRepo)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	_productHandler.NewProductHandler(e, productUsecase)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Data.AppPort)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
