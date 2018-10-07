package app

import (
	"net/http"
	"sync/atomic"

	"github.com/ilovelili/simplewebserver/util"
)

// ops atomic counter
var ops uint64

// Root root
func (app *App) Root(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, "root")
}

// GetCounter get counter
func (app *App) GetCounter(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, struct {
		Count uint64 `json:"count"`
	}{ops})
}

// SetCounter set counter
func (app *App) SetCounter(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&ops, 1)
	app.GetCounter(w, r)
}
