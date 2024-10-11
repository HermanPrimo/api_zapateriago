package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HermanPrimo/api_zapateriago/data"
	"github.com/HermanPrimo/api_zapateriago/models"
	"github.com/HermanPrimo/api_zapateriago/utils"
	"github.com/gorilla/mux"
)

// GetGeneros obtiene la lista de Generos.
func GetGeneros(w http.ResponseWriter, r *http.Request) {
	var Generos []models.Genero
	if result := data.DB.Find(&Generos); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de Generos",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de Generos",
		StatusCode: http.StatusOK,
		Data:       Generos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// GetGenero obtiene un Genero por su ID.
func GetGenero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var genero models.Genero

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

	if result := data.DB.Where("gen_id = ?", id).First(&genero); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Genero no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Genero",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Genero encontrado",
		StatusCode: http.StatusOK,
		Data:       genero,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// NewGenero crea un nuevo Genero.
func NewGenero(w http.ResponseWriter, r *http.Request) {
	var genero models.Genero
	if err := json.NewDecoder(r.Body).Decode(&genero); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&genero); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el Genero",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Genero creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       genero,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// DeleteGenero elimina un Genero por su ID.
func DeleteGenero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var genero models.Genero

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

	if result := data.DB.Where("gen_id = ?", id).Delete(&genero); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Genero no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el Genero",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Genero eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// UpdateGenero actualiza un Genero existente.
func UpdateGenero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var genero models.Genero

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
	if result := data.DB.Where("gen_id = ?", id).First(&genero); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Genero no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Genero",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updatedGenero models.Genero
	if err := json.NewDecoder(r.Body).Decode(&updatedGenero); err != nil {
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
	data.DB.Model(&genero).Updates(updatedGenero)

	respuesta := utils.Respuesta{
		Msg:        "Genero actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       genero,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
