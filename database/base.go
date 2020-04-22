package database

import (
	"fmt"
	"log"
	"os"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init se conecta a la BD y crea la tabla "contacts"
func Init() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DB")
	dbHost := os.Getenv("MYSQL_HOST")
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName) // Formatear un string con parámetros
	log.Printf(dbURI)                                                                                               // Imprimir la URI de conexión para debug

	conn, err := gorm.Open("mysql", dbURI) // Abrir la conexión con MySQL
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Contact{}) // Automáticamente generar tabla para el modelo
}

// DB regresa el objeto de base de datos
func DB() *gorm.DB {
	return db
}
