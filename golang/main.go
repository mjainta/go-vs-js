package main

import (
	"github.com/mjainta/go-vs-js/app"
)


func main() {
	r := app.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
