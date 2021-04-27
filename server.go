package main

import (
	route "echo-template/route"
)

func main() {
	//Routes
	e := route.Init()

	// Start local server
	e.Logger.Fatal(e.Start(":80"))
}
