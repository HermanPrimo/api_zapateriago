package data

import (
	"log"
	"os"

	// Asegúrate de importar tu paquete de modelos
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarPostgres() {

	var error error
	DB, error = gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if error != nil {
		log.Fatal("Error al conectar a la base de datos:", error)
	} else {
		log.Println("Conectado a la Base de Datos")
	}

}

/*

// Sincronizar la base de datos con el modelo
	if err := DB.AutoMigrate(&models.Rol{}); err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}

package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION_STRING = "host=localhost user=postgres password=190818 dbname=prueba port=5432 sslmode=disable TimeZone=America/Mexico_City"
var DB *gorm.DB

func ConectarPostgres() {
	var error error
	DB, error = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})
	if error != nil {
		log.Fatal("Error al concectar la base de datos:", error)
	} else {
		log.Println("Conectado a la Base de Datos")
	}
}
*/
