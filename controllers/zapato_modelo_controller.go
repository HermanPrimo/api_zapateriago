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

// GetZapatoModelos obtiene la lista de ZapatoModelos.
func GetZapatoModelos(w http.ResponseWriter, r *http.Request) {
	var ZapatoModelos []models.ZapatoModelo
	if result := data.DB.Find(&ZapatoModelos); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de ZapatoModelos",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de ZapatoModelos",
		StatusCode: http.StatusOK,
		Data:       ZapatoModelos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// GetZapatoModelo obtiene un ZapatoModelo por su ID.
func GetZapatoModelo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoModelo models.ZapatoModelo

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

	if result := data.DB.Where("zapmod_id = ?", id).First(&zapatoModelo); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoModelo no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el ZapatoModelo",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoModelo encontrado",
		StatusCode: http.StatusOK,
		Data:       zapatoModelo,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// NewZapatoModelo crea un nuevo ZapatoModelo.
func NewZapatoModelo(w http.ResponseWriter, r *http.Request) {
	var zapatoModelo models.ZapatoModelo
	if err := json.NewDecoder(r.Body).Decode(&zapatoModelo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&zapatoModelo); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el ZapatoModelo",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoModelo creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       zapatoModelo,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// DeleteZapatoModelo elimina un ZapatoModelo por su ID.
func DeleteZapatoModelo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoModelo models.ZapatoModelo

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

	if result := data.DB.Where("zapmod_id = ?", id).Delete(&zapatoModelo); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoModelo no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el ZapatoModelo",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoModelo eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// UpdateZapatoModelo actualiza un ZapatoModelo existente.
func UpdateZapatoModelo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoModelo models.ZapatoModelo

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
	if result := data.DB.Where("zapmod_id = ?", id).First(&zapatoModelo); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoModelo no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el ZapatoModelo",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updatedZapatoModelo models.ZapatoModelo
	if err := json.NewDecoder(r.Body).Decode(&updatedZapatoModelo); err != nil {
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
	data.DB.Model(&zapatoModelo).Updates(updatedZapatoModelo)

	respuesta := utils.Respuesta{
		Msg:        "ZapatoModelo actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       zapatoModelo,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
