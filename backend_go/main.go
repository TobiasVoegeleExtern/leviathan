package main

import (
	"backend_go/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8000") // Run the backend on port 8080
}
