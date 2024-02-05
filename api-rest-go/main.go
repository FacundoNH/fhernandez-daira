package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Operación representa una operación matemática.
type Operacion struct {
	Num1      int    `json:"num1"`
	Num2      int    `json:"num2"`
	Operador  string `json:"operador"`
	Resultado int    `json:"resultado"`
}

var historial []Operacion

func main() {
	r := mux.NewRouter()

	// Rutas para realizar operaciones y manejar el historial
	r.HandleFunc("/api/operar", Operar).Methods("POST")
	r.HandleFunc("/api/historial", ObtenerHistorial).Methods("GET")

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// Operar realiza operaciones matemáticas y almacena el resultado en el historial.
func Operar(w http.ResponseWriter, r *http.Request) {
	var operacion Operacion
	err := json.NewDecoder(r.Body).Decode(&operacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validar entrada
	if operacion.Num1 < 0 || operacion.Num2 < 0 {
		http.Error(w, "Los números deben ser positivos", http.StatusBadRequest)
		return
	}

	// Realizar operación
	switch operacion.Operador {
	case "+":
		operacion.Resultado = operacion.Num1 + operacion.Num2
	case "-":
		operacion.Resultado = operacion.Num1 - operacion.Num2
	case "*":
		operacion.Resultado = operacion.Num1 * operacion.Num2
	case "/":
		if operacion.Num2 != 0 {
			operacion.Resultado = operacion.Num1 / operacion.Num2
		} else {
			http.Error(w, "No se puede dividir por cero", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Operador no válido", http.StatusBadRequest)
		return
	}

	// Agregar la operación al historial
	historial = append(historial, operacion)

	// Responder con el resultado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operacion)
}

// ObtenerHistorial devuelve el historial de operaciones.
func ObtenerHistorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historial)
}
