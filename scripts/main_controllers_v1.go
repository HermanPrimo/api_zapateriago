package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const controllerTemplate = `package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HermanPrimo/api_zapateriago/data"
	"github.com/HermanPrimo/api_zapateriago/models"
	"github.com/HermanPrimo/api_zapateriago/utils"
	"github.com/gorilla/mux"
)

// Get{{.ModelName}}s obtiene la lista de {{.ModelName}}s.
func Get{{.ModelName}}s(w http.ResponseWriter, r *http.Request) {
	var {{.ModelPlural}} []models.{{.ModelName}}
	if result := data.DB.Find(&{{.ModelPlural}}); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de {{.ModelPlural}}",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de {{.ModelPlural}}",
		StatusCode: http.StatusOK,
		Data:       {{.ModelPlural}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Get{{.ModelName}} obtiene un {{.ModelName}} por su ID.
func Get{{.ModelName}}(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var {{.ModelVar}} models.{{.ModelName}}

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "ID inválido",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Where("{{.IDColumn}} = ?", id).First(&{{.ModelVar}}); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "{{.ModelName}} no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el {{.ModelName}}",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "{{.ModelName}} encontrado",
		StatusCode: http.StatusOK,
		Data:       {{.ModelVar}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// New{{.ModelName}} crea un nuevo {{.ModelName}}.
func New{{.ModelName}}(w http.ResponseWriter, r *http.Request) {
	var {{.ModelVar}} models.{{.ModelName}}
	if err := json.NewDecoder(r.Body).Decode(&{{.ModelVar}}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&{{.ModelVar}}); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el {{.ModelName}}",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "{{.ModelName}} creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       {{.ModelVar}},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// Delete{{.ModelName}} elimina un {{.ModelName}} por su ID.
func Delete{{.ModelName}}(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var {{.ModelVar}} models.{{.ModelName}}

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "ID inválido",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Where("{{.IDColumn}} = ?", id).Delete(&{{.ModelVar}}); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "{{.ModelName}} no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el {{.ModelName}}",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "{{.ModelName}} eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Update{{.ModelName}} actualiza un {{.ModelName}} existente.
func Update{{.ModelName}}(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var {{.ModelVar}} models.{{.ModelName}}

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "ID inválido",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	// Buscar el registro existente
	if result := data.DB.Where("{{.IDColumn}} = ?", id).First(&{{.ModelVar}}); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "{{.ModelName}} no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el {{.ModelName}}",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updated{{.ModelName}} models.{{.ModelName}}
	if err := json.NewDecoder(r.Body).Decode(&updated{{.ModelName}}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	// Actualizar solo los campos necesarios del modelo existente
	data.DB.Model(&{{.ModelVar}}).Updates(updated{{.ModelName}})

	respuesta := utils.Respuesta{
		Msg:        "{{.ModelName}} actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       {{.ModelVar}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
`

func main() {
	modelDir := "./models"
	files, err := os.ReadDir(modelDir)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			filePath := filepath.Join(modelDir, file.Name())
			processModelFile(filePath)
		}
	}
}

func processModelFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	var modelName string
	var modelPlural string
	var idColumn string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "type ") {
			modelName = strings.Split(line, " ")[1]
		}

		if strings.Contains(line, "gorm:\"column:") && strings.Contains(line, "primaryKey") {
			idColumn = extractColumnName(line)
		}
	}

	modelVar := strings.ToLower(string(modelName[0])) + modelName[1:]
	modelPlural = modelName + "s"

	data := struct {
		ModelName   string
		ModelVar    string
		ModelPlural string
		IDColumn    string
	}{
		ModelName:   modelName,
		ModelVar:    modelVar,
		ModelPlural: modelPlural,
		IDColumn:    idColumn,
	}

	tmpl, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		fmt.Println("Error al parsear la plantilla:", err)
		return
	}

	outputFilePath := filepath.Join("./controllers", fmt.Sprintf("%s_controller.go", modelVar))
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	err = tmpl.Execute(writer, data)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla:", err)
		return
	}

	writer.Flush()
}

func extractColumnName(line string) string {
	parts := strings.Split(line, "gorm:\"column:")
	columnPart := strings.Split(parts[1], ";")
	return columnPart[0]
}
