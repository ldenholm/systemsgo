package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ldenholm/systemsgo/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	// create servemux
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)
	http.ListenAndServe(":9090", sm)
}
