package main

import (
	route "ksd-grm-api/route"
)

func main() {
	//Routes
	e := route.Init()

	// Start local server
	e.Logger.Fatal(e.Start(":80"))
}
