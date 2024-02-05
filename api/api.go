// api/api.go
package api

import (
	"encoding/json"
	"net/http"
	backend "proyecto-calculadora/BackEnd"

	"github.com/gorilla/mux"
)

// IniciarAPI inicializa la API.
func IniciarAPI() {
	r := mux.NewRouter()

	// Rutas para realizar operaciones y manejar el historial
	r.HandleFunc("/api/operar", Operar).Methods("POST")
	r.HandleFunc("/api/historial", ObtenerHistorial).Methods("GET")

	http.ListenAndServe(":8080", r)
}

// Operar realiza operaciones matemáticas y almacena el resultado en el historial.
func Operar(w http.ResponseWriter, r *http.Request) {
	var operacion backend.Operacion
	err := json.NewDecoder(r.Body).Decode(&operacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := backend.RealizarOperacion(operacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Responder con el resultado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"resultado": resultado})
}

// ObtenerHistorial devuelve el historial de operaciones.
func ObtenerHistorial(w http.ResponseWriter, r *http.Request) {
	historial := backend.ObtenerHistorial()

	// Responder con el historial
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historial)
}
