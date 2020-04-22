package main

import (
	"directorio-tap/controllers"
	"directorio-tap/database"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CustomValidator is a custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates using a CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Load config from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err) // Imprimir en consola y terminar el programa.
	}

	database.Init()

	e := echo.New()
	e.Use(middleware.Logger()) // marcará error en VS Code, pero no importa. :3
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/contact", controllers.CreateContact)
	e.GET("/contact", controllers.GetContacts)
	e.GET("/contact/:id", controllers.GetContact)
	e.PUT("/contact/:id", controllers.UpdateContact)
	e.DELETE("/contact/:id", controllers.DeleteContact)

	// Start Echo HTTP server
	e.Logger.Fatal(e.Start(":" + os.Getenv("HTTP_PORT"))) // Iniciar servidor en el puerto definido.
}
