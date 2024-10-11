package routes

import (
	"github.com/HermanPrimo/api_zapateriago/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	// Rutas para Categoria
	apiCategoria := api.PathPrefix("/categoria").Subrouter()
	apiCategoria.HandleFunc("", controllers.GetCategorias).Methods("GET")     // Plural para obtener todos
	apiCategoria.HandleFunc("/{id}", controllers.GetCategoria).Methods("GET") // Singular para obtener uno
	apiCategoria.HandleFunc("", controllers.NewCategoria).Methods("POST")
	apiCategoria.HandleFunc("/{id}", controllers.DeleteCategoria).Methods("DELETE")
	apiCategoria.HandleFunc("/{id}", controllers.UpdateCategoria).Methods("PUT")

	// Rutas para Color
	apiColor := api.PathPrefix("/color").Subrouter()
	apiColor.HandleFunc("", controllers.GetColors).Methods("GET")     // Plural para obtener todos
	apiColor.HandleFunc("/{id}", controllers.GetColor).Methods("GET") // Singular para obtener uno
	apiColor.HandleFunc("", controllers.NewColor).Methods("POST")
	apiColor.HandleFunc("/{id}", controllers.DeleteColor).Methods("DELETE")
	apiColor.HandleFunc("/{id}", controllers.UpdateColor).Methods("PUT")

	// Rutas para DetalleVenta
	apiDetalleVenta := api.PathPrefix("/detalle-venta").Subrouter()
	apiDetalleVenta.HandleFunc("", controllers.GetDetalleVentas).Methods("GET")     // Plural para obtener todos
	apiDetalleVenta.HandleFunc("/{id}", controllers.GetDetalleVenta).Methods("GET") // Singular para obtener uno
	apiDetalleVenta.HandleFunc("", controllers.NewDetalleVenta).Methods("POST")
	apiDetalleVenta.HandleFunc("/{id}", controllers.DeleteDetalleVenta).Methods("DELETE")
	apiDetalleVenta.HandleFunc("/{id}", controllers.UpdateDetalleVenta).Methods("PUT")

	// Rutas para Empleado
	apiEmpleado := api.PathPrefix("/empleado").Subrouter()
	apiEmpleado.HandleFunc("", controllers.GetEmpleados).Methods("GET")     // Plural para obtener todos
	apiEmpleado.HandleFunc("/{id}", controllers.GetEmpleado).Methods("GET") // Singular para obtener uno
	apiEmpleado.HandleFunc("", controllers.NewEmpleado).Methods("POST")
	apiEmpleado.HandleFunc("/{id}", controllers.DeleteEmpleado).Methods("DELETE")
	apiEmpleado.HandleFunc("/{id}", controllers.UpdateEmpleado).Methods("PUT")

	// Rutas para Estado
	apiEstado := api.PathPrefix("/estado").Subrouter()
	apiEstado.HandleFunc("", controllers.GetEstados).Methods("GET")     // Plural para obtener todos
	apiEstado.HandleFunc("/{id}", controllers.GetEstado).Methods("GET") // Singular para obtener uno
	apiEstado.HandleFunc("", controllers.NewEstado).Methods("POST")
	apiEstado.HandleFunc("/{id}", controllers.DeleteEstado).Methods("DELETE")
	apiEstado.HandleFunc("/{id}", controllers.UpdateEstado).Methods("PUT")

	// Rutas para Genero
	apiGenero := api.PathPrefix("/genero").Subrouter()
	apiGenero.HandleFunc("", controllers.GetGeneros).Methods("GET")     // Plural para obtener todos
	apiGenero.HandleFunc("/{id}", controllers.GetGenero).Methods("GET") // Singular para obtener uno
	apiGenero.HandleFunc("", controllers.NewGenero).Methods("POST")
	apiGenero.HandleFunc("/{id}", controllers.DeleteGenero).Methods("DELETE")
	apiGenero.HandleFunc("/{id}", controllers.UpdateGenero).Methods("PUT")

	// Rutas para Marca
	apiMarca := api.PathPrefix("/marca").Subrouter()
	apiMarca.HandleFunc("", controllers.GetMarcas).Methods("GET")     // Plural para obtener todos
	apiMarca.HandleFunc("/{id}", controllers.GetMarca).Methods("GET") // Singular para obtener uno
	apiMarca.HandleFunc("", controllers.NewMarca).Methods("POST")
	apiMarca.HandleFunc("/{id}", controllers.DeleteMarca).Methods("DELETE")
	apiMarca.HandleFunc("/{id}", controllers.UpdateMarca).Methods("PUT")

	// Rutas para Modelo
	apiModelo := api.PathPrefix("/modelo").Subrouter()
	apiModelo.HandleFunc("", controllers.GetModelos).Methods("GET")     // Plural para obtener todos
	apiModelo.HandleFunc("/{id}", controllers.GetModelo).Methods("GET") // Singular para obtener uno
	apiModelo.HandleFunc("", controllers.NewModelo).Methods("POST")
	apiModelo.HandleFunc("/{id}", controllers.DeleteModelo).Methods("DELETE")
	apiModelo.HandleFunc("/{id}", controllers.UpdateModelo).Methods("PUT")

	// Rutas para Pago
	apiPago := api.PathPrefix("/pago").Subrouter()
	apiPago.HandleFunc("", controllers.GetPagos).Methods("GET")     // Plural para obtener todos
	apiPago.HandleFunc("/{id}", controllers.GetPago).Methods("GET") // Singular para obtener uno
	apiPago.HandleFunc("", controllers.NewPago).Methods("POST")
	apiPago.HandleFunc("/{id}", controllers.DeletePago).Methods("DELETE")
	apiPago.HandleFunc("/{id}", controllers.UpdatePago).Methods("PUT")

	// Rutas para Puesto
	apiPuesto := api.PathPrefix("/puesto").Subrouter()
	apiPuesto.HandleFunc("", controllers.GetPuestos).Methods("GET")     // Plural para obtener todos
	apiPuesto.HandleFunc("/{id}", controllers.GetPuesto).Methods("GET") // Singular para obtener uno
	apiPuesto.HandleFunc("", controllers.NewPuesto).Methods("POST")
	apiPuesto.HandleFunc("/{id}", controllers.DeletePuesto).Methods("DELETE")
	apiPuesto.HandleFunc("/{id}", controllers.UpdatePuesto).Methods("PUT")

	// Rutas para Talla
	apiTalla := api.PathPrefix("/talla").Subrouter()
	apiTalla.HandleFunc("", controllers.GetTallas).Methods("GET")     // Plural para obtener todos
	apiTalla.HandleFunc("/{id}", controllers.GetTalla).Methods("GET") // Singular para obtener uno
	apiTalla.HandleFunc("", controllers.NewTalla).Methods("POST")
	apiTalla.HandleFunc("/{id}", controllers.DeleteTalla).Methods("DELETE")
	apiTalla.HandleFunc("/{id}", controllers.UpdateTalla).Methods("PUT")

	// Rutas para Tipo
	apiTipo := api.PathPrefix("/tipo").Subrouter()
	apiTipo.HandleFunc("", controllers.GetTipos).Methods("GET")     // Plural para obtener todos
	apiTipo.HandleFunc("/{id}", controllers.GetTipo).Methods("GET") // Singular para obtener uno
	apiTipo.HandleFunc("", controllers.NewTipo).Methods("POST")
	apiTipo.HandleFunc("/{id}", controllers.DeleteTipo).Methods("DELETE")
	apiTipo.HandleFunc("/{id}", controllers.UpdateTipo).Methods("PUT")

	// Rutas para Usuario
	apiUsuario := api.PathPrefix("/usuario").Subrouter()
	apiUsuario.HandleFunc("", controllers.GetUsuarios).Methods("GET")     // Plural para obtener todos
	apiUsuario.HandleFunc("/{id}", controllers.GetUsuario).Methods("GET") // Singular para obtener uno
	apiUsuario.HandleFunc("", controllers.NewUsuario).Methods("POST")
	apiUsuario.HandleFunc("/{id}", controllers.DeleteUsuario).Methods("DELETE")
	apiUsuario.HandleFunc("/{id}", controllers.UpdateUsuario).Methods("PUT")

	// Rutas para Venta
	apiVenta := api.PathPrefix("/venta").Subrouter()
	apiVenta.HandleFunc("", controllers.GetVentas).Methods("GET")     // Plural para obtener todos
	apiVenta.HandleFunc("/{id}", controllers.GetVenta).Methods("GET") // Singular para obtener uno
	apiVenta.HandleFunc("", controllers.NewVenta).Methods("POST")
	apiVenta.HandleFunc("/{id}", controllers.DeleteVenta).Methods("DELETE")
	apiVenta.HandleFunc("/{id}", controllers.UpdateVenta).Methods("PUT")

	// Rutas para Zapato
	apiZapato := api.PathPrefix("/zapato").Subrouter()
	apiZapato.HandleFunc("", controllers.GetZapatos).Methods("GET")     // Plural para obtener todos
	apiZapato.HandleFunc("/{id}", controllers.GetZapato).Methods("GET") // Singular para obtener uno
	apiZapato.HandleFunc("", controllers.NewZapato).Methods("POST")
	apiZapato.HandleFunc("/{id}", controllers.DeleteZapato).Methods("DELETE")
	apiZapato.HandleFunc("/{id}", controllers.UpdateZapato).Methods("PUT")

	// Rutas para ZapatoModelo
	apiZapatoModelo := api.PathPrefix("/zapato-modelo").Subrouter()
	apiZapatoModelo.HandleFunc("", controllers.GetZapatoModelos).Methods("GET")     // Plural para obtener todos
	apiZapatoModelo.HandleFunc("/{id}", controllers.GetZapatoModelo).Methods("GET") // Singular para obtener uno
	apiZapatoModelo.HandleFunc("", controllers.NewZapatoModelo).Methods("POST")
	apiZapatoModelo.HandleFunc("/{id}", controllers.DeleteZapatoModelo).Methods("DELETE")
	apiZapatoModelo.HandleFunc("/{id}", controllers.UpdateZapatoModelo).Methods("PUT")

	// Rutas para ZapatoTalla
	apiZapatoTalla := api.PathPrefix("/zapato-talla").Subrouter()
	apiZapatoTalla.HandleFunc("", controllers.GetZapatoTallas).Methods("GET")     // Plural para obtener todos
	apiZapatoTalla.HandleFunc("/{id}", controllers.GetZapatoTalla).Methods("GET") // Singular para obtener uno
	apiZapatoTalla.HandleFunc("", controllers.NewZapatoTalla).Methods("POST")
	apiZapatoTalla.HandleFunc("/{id}", controllers.DeleteZapatoTalla).Methods("DELETE")
	apiZapatoTalla.HandleFunc("/{id}", controllers.UpdateZapatoTalla).Methods("PUT")

	return rutas
}
