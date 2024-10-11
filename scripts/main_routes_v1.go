package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Ruta a la carpeta de modelos
	modelosPath := "./models"
	var modelos []string

	// Buscar los archivos Go en la carpeta de modelos
	err := filepath.Walk(modelosPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Solo procesa archivos .go
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			// Abrir y leer el archivo
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Usar un escáner para leer el archivo línea por línea
			scanner := bufio.NewScanner(file)
			structRegex := regexp.MustCompile(`type (\w+) struct`) // Regex para encontrar structs

			for scanner.Scan() {
				line := scanner.Text()
				// Buscar estructuras (structs)
				match := structRegex.FindStringSubmatch(line)
				if len(match) > 1 {
					modelos = append(modelos, match[1])
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al leer la carpeta de modelos:", err)
		return
	}

	// Genera el archivo index_route.go en la carpeta ./routes
	filePath := "./routes/index_route.go"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el encabezado del archivo con los paquetes
	file.WriteString(`package routes

import (
	"github.com/HermanPrimo/api_zapateriago/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

`)

	// Generar rutas para cada modelo
	for _, modelo := range modelos {
		modeloLower := strings.ToLower(modelo)
		ruta := fmt.Sprintf(`
	// Rutas para %s
	api%s := api.PathPrefix("/%s").Subrouter()
	api%s.HandleFunc("", controllers.Get%ss).Methods("GET")  // Plural para obtener todos
	api%s.HandleFunc("/{id}", controllers.Get%s).Methods("GET") // Singular para obtener uno
	api%s.HandleFunc("", controllers.New%s).Methods("POST")
	api%s.HandleFunc("/{id}", controllers.Delete%s).Methods("DELETE")
	api%s.HandleFunc("/{id}", controllers.Update%s).Methods("PUT")
`, modelo, modelo, modeloLower, modelo, modelo, modelo, modelo, modelo, modelo, modelo, modelo, modelo)

		file.WriteString(ruta)
	}

	// Cerrar la función InitRouter
	file.WriteString(`
	return rutas
}
`)

	fmt.Println("Rutas generadas correctamente en", filePath)
}
