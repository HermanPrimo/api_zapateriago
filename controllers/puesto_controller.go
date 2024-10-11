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

// GetPuestos obtiene la lista de Puestos.
func GetPuestos(w http.ResponseWriter, r *http.Request) {
	var Puestos []models.Puesto
	if result := data.DB.Find(&Puestos); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de Puestos",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de Puestos",
		StatusCode: http.StatusOK,
		Data:       Puestos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// GetPuesto obtiene un Puesto por su ID.
func GetPuesto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var puesto models.Puesto

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

	if result := data.DB.Where("pue_id = ?", id).First(&puesto); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Puesto no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Puesto",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Puesto encontrado",
		StatusCode: http.StatusOK,
		Data:       puesto,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// NewPuesto crea un nuevo Puesto.
func NewPuesto(w http.ResponseWriter, r *http.Request) {
	var puesto models.Puesto
	if err := json.NewDecoder(r.Body).Decode(&puesto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&puesto); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el Puesto",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Puesto creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       puesto,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// DeletePuesto elimina un Puesto por su ID.
func DeletePuesto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var puesto models.Puesto

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

	if result := data.DB.Where("pue_id = ?", id).Delete(&puesto); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Puesto no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el Puesto",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Puesto eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// UpdatePuesto actualiza un Puesto existente.
func UpdatePuesto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var puesto models.Puesto

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
	if result := data.DB.Where("pue_id = ?", id).First(&puesto); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "Puesto no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el Puesto",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updatedPuesto models.Puesto
	if err := json.NewDecoder(r.Body).Decode(&updatedPuesto); err != nil {
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
	data.DB.Model(&puesto).Updates(updatedPuesto)

	respuesta := utils.Respuesta{
		Msg:        "Puesto actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       puesto,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
