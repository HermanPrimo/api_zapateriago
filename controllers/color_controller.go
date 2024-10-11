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

// GetColors obtiene la lista de Colors.
func GetColors(w http.ResponseWriter, r *http.Request) {
	var Colors []models.Color
	if result := data.DB.Find(&Colors); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de Colors",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de Colors",
		StatusCode: http.StatusOK,
		Data:       Colors,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// GetColor obtiene un Color por su ID.
func GetColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var color models.Color

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

	if result := data.DB.Where("col_id = ?", id).First(&color); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Color no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Color",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Color encontrado",
		StatusCode: http.StatusOK,
		Data:       color,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// NewColor crea un nuevo Color.
func NewColor(w http.ResponseWriter, r *http.Request) {
	var color models.Color
	if err := json.NewDecoder(r.Body).Decode(&color); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&color); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el Color",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Color creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       color,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// DeleteColor elimina un Color por su ID.
func DeleteColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var color models.Color

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

	if result := data.DB.Where("col_id = ?", id).Delete(&color); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Color no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el Color",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Color eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// UpdateColor actualiza un Color existente.
func UpdateColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var color models.Color

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
	if result := data.DB.Where("col_id = ?", id).First(&color); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Color no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Color",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updatedColor models.Color
	if err := json.NewDecoder(r.Body).Decode(&updatedColor); err != nil {
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
	data.DB.Model(&color).Updates(updatedColor)

	respuesta := utils.Respuesta{
		Msg:        "Color actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       color,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
