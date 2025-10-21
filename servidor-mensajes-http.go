package main

//SERVIDOR MENSAJES
import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mensajes []string
	mu       sync.Mutex
)

func handleInicio(w http.ResponseWriter, r *http.Request) {
	mensaje := "Bienvenidos al servidor de mensajes HTTP \n"
	mensaje += "Rutas : \n"
	mensaje += " /guardar?mensaje=texto \n"
	mensaje += " /mensajes \n"
	fmt.Fprintf(w, mensaje)
}

func handleGuardar(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("mensaje")
	if msg == "" {
		fmt.Fprintf(w, "ERROR MENSAJE VACIO")
		return
	}
	mu.Lock()
	mensajes = append(mensajes, msg)
	mu.Unlock()
	fmt.Fprintf(w, "Mensaje Guardado")
}

func handleMensajes(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	if len(mensajes) == 0 {
		fmt.Fprintf(w, "No hay mensajes")
	}
	for i, msg := range mensajes {
		fmt.Fprintf(w, "%d: %s \n", i+1, msg)
	}

}

func main() {
	http.HandleFunc("/", handleInicio)
	http.HandleFunc("/guardar", handleGuardar)
	http.HandleFunc("/listarmensajes", handleMensajes)

	fmt.Println("Servidor HTTP en http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
