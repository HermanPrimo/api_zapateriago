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

// GetZapatoTallas obtiene la lista de ZapatoTallas.
func GetZapatoTallas(w http.ResponseWriter, r *http.Request) {
	var ZapatoTallas []models.ZapatoTalla
	if result := data.DB.Find(&ZapatoTallas); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al obtener la lista de ZapatoTallas",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Lista de ZapatoTallas",
		StatusCode: http.StatusOK,
		Data:       ZapatoTallas,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// GetZapatoTalla obtiene un ZapatoTalla por su ID.
func GetZapatoTalla(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoTalla models.ZapatoTalla

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

	if result := data.DB.Where("zaptal_id = ?", id).First(&zapatoTalla); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoTalla no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el ZapatoTalla",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoTalla encontrado",
		StatusCode: http.StatusOK,
		Data:       zapatoTalla,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// NewZapatoTalla crea un nuevo ZapatoTalla.
func NewZapatoTalla(w http.ResponseWriter, r *http.Request) {
	var zapatoTalla models.ZapatoTalla
	if err := json.NewDecoder(r.Body).Decode(&zapatoTalla); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Msg:        "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if result := data.DB.Create(&zapatoTalla); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al crear el ZapatoTalla",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoTalla creado exitosamente",
		StatusCode: http.StatusCreated,
		Data:       zapatoTalla,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// DeleteZapatoTalla elimina un ZapatoTalla por su ID.
func DeleteZapatoTalla(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoTalla models.ZapatoTalla

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

	if result := data.DB.Where("zaptal_id = ?", id).Delete(&zapatoTalla); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoTalla no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al eliminar el ZapatoTalla",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "ZapatoTalla eliminado correctamente",
		StatusCode: http.StatusOK,
		Data:       nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// UpdateZapatoTalla actualiza un ZapatoTalla existente.
func UpdateZapatoTalla(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var zapatoTalla models.ZapatoTalla

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
	if result := data.DB.Where("zaptal_id = ?", id).First(&zapatoTalla); result.Error != nil {
		if result.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			respuesta := utils.Respuesta{
				Msg:        "ZapatoTalla no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Msg:        "Error al buscar el ZapatoTalla",
			StatusCode: http.StatusInternalServerError,
			Data:       result.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	var updatedZapatoTalla models.ZapatoTalla
	if err := json.NewDecoder(r.Body).Decode(&updatedZapatoTalla); err != nil {
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
	data.DB.Model(&zapatoTalla).Updates(updatedZapatoTalla)

	respuesta := utils.Respuesta{
		Msg:        "ZapatoTalla actualizado exitosamente",
		StatusCode: http.StatusOK,
		Data:       zapatoTalla,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}
