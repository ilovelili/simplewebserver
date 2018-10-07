// Package app application container and routing config module
package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// App app object
type App struct {
	Router         *mux.Router
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Serve serve Non-TLS with cros origin support
func (app *App) Serve(addr string) {
	handler := cors.Default().Handler(app.Router)
	fmt.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

// Initialize init the app
func (app *App) Initialize() {
	// set up new router
	app.Router = mux.NewRouter()
	// init routes
	app.InitializeRoutes()
}
