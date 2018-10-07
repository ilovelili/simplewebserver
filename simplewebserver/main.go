package main

import (
	"github.com/ilovelili/simplewebserver/app"
)

// main Entry
func main() {
	app := &app.App{}
	app.Initialize()
	app.Serve(":8080")
}
