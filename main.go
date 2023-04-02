package main

import (
	"log"
)

func main() {
	r := BootstrapRouter()
	err := r.Run()
	if err != nil {
		log.Fatal("Unable to create the gin server")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
