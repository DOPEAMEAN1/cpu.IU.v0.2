import (
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

package main

	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

var log = logrus.New()

func main() {
	http.HandleFunc("/", indexHandler)
	port := viper.GetString("PORT")
	if port == "" {
		port = defaultPort
		log.Infof("Defaulting to port %s", port)
	}

	log.Infof("Listening on port %s", port)
	log.Infof("Open http://localhost:%s in the browser", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleNotFoundError(w, r)
		return
	}

	tmpl, err := template.New("index").Parse(`<html><body>Hello, World!</body></html>`)
	if err != nil {
		handleInternalServerError(w, r, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		handleInternalServerError(w, r, err)
	}
}

func handleNotFoundError(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	// Additional error handling logic can be added here
}

func handleInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	// Additional error handling logic can be added here
}

func startServer(port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform logging logic here
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}