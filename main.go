package main

import (
	// // open comment to deploy gcp app engine
	// "google.golang.org/appengine"

	"github.com/jojoarianto/go-ddd-api/interfaces"
)

func main() {
	// // open comment to deploy gcp app engine
	// router := interfaces.Routes()
	// http.Handle("/", router)
	// appengine.Main() // Start the gcp server

	interfaces.Run(8000)
}
