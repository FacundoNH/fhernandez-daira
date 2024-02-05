package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Operacion representa una operación matemática.
type Operacion struct {
	Num1      int    `json:"num1"`
	Num2      int    `json:"num2"`
	Operador  string `json:"operador"`
	Resultado int    `json:"resultado"`
}

var historialMutex sync.Mutex
var historial []Operacion

func IniciarBackend() {
	// Puedes inicializar cualquier configuración necesaria para el backend aquí
	// Por ejemplo, podrías inicializar la conexión a la base de datos, etc.
	fmt.Println("Backend iniciado")
}

func RealizarOperacion(operacion Operacion) (int, error) {
	historialMutex.Lock()
	defer historialMutex.Unlock()

	// Validar entrada
	if operacion.Num1 < 0 || operacion.Num2 < 0 {
		return 0, fmt.Errorf("Los números deben ser positivos")
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
			return 0, fmt.Errorf("No se puede dividir por cero")
		}
	default:
		return 0, fmt.Errorf("Operador no válido")
	}

	// Agregar la operación al historial
	historial = append(historial, operacion)

	return operacion.Resultado, nil
}

func ObtenerHistorial() []Operacion {
	historialMutex.Lock()
	defer historialMutex.Unlock()

	return historial
}

// Manejador para la API de operaciones
func OperarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el cuerpo JSON
	var operacion Operacion
	err := json.NewDecoder(r.Body).Decode(&operacion)
	if err != nil {
		http.Error(w, "Error al decodificar la operación JSON", http.StatusBadRequest)
		return
	}

	// Realizar la operación
	resultado, err := RealizarOperacion(operacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Responder con el resultado JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"resultado": resultado})
}

// Manejador para obtener el historial de operaciones
func HistorialHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el historial de operaciones
	historial := ObtenerHistorial()

	// Responder con el historial JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historial)
}
