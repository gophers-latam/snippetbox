package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	} 
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extraer el valor del parámetro id de la cadena de consulta y tratar de
	// convertirlo a un entero usando la función strconv.Atoi(). Si no puede
	// convertirse a un entero o el valor es menor que 1, devolvemos una página 404
	// no encontrada.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Utilizar la función fmt.Fprintf() para interpolar el valor id con nuestra respuesta
	// y escribirlo en http.ResponseWriter.
	fmt.Fprintf(w, "Mostrar un fragmento específico con ID %d...", id)
}


func prueba(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("desde prueba"))
}

// Agrega una función de controlador snippetCreate.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Si no lo es, usa el método w.WriteHeader() para enviar un código de estado 405
		// y el método w.Write() para escribir un cuerpo de respuesta "Method Not Allowed".
		// Luego, retornamos de la función para que el código subsiguiente no se ejecute.
		w.Header().Set("Allow", "POST") 
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)


		// WriteHeader sends an HTTP response header with the provided
		// status code.
		//
		// If WriteHeader is not called explicitly, the first call to Write
		// will trigger an implicit WriteHeader(http.StatusOK).
		// Thus explicit calls to WriteHeader are mainly used to
		// send error codes or 1xx informational responses.
		//
		// The provided code must be a valid HTTP 1xx-5xx status code.
		// Any number of 1xx headers may be written, followed by at most
		// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
		// headers may be buffered. Use the Flusher interface to send
		// buffered data. The header map is cleared when 2xx-5xx headers are
		// sent, but not with 1xx headers.
		//
		// The server will automatically send a 100 (Continue) header
		// on the first read from the request body if the request has
		// an "Expect: 100-continue" header.
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	} 
	w.Write([]byte("Crear un nuevo fragmento..."))
}

func main() {
	// Registra las dos nuevas funciones de controlador y los patrones de URL correspondientes
	// con el servemux, de la misma manera que hicimos antes.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/prueba", prueba)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
