// Package app application container and routing config module
package app

// InitializeRoutes init routes
func (app *App) InitializeRoutes() {
	// show root
	app.Router.HandleFunc("/", app.Root).Methods("GET", "POST")
	// get counter
	app.Router.HandleFunc("/count", app.GetCounter).Methods("GET")
	// set counter
	app.Router.HandleFunc("/count", app.SetCounter).Methods("POST")
}
