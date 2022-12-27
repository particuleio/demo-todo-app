package main

import (
	"fmt"
	"log"
	"net/http"
)

type routeMap map[string]func(w http.ResponseWriter, r *http.Request)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Called function [ping]")
	fmt.Fprintf(w, "pong")
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("Called function [create]")
	fmt.Fprintf(w, "pong")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Called function [getAll]")
	fmt.Fprintf(w, "pong")
}

func delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Called function [delete]")
	fmt.Fprintf(w, "pong")
}

func main() {

	rm := routeMap{
		"/ping":   ping,
		"/create": create,
		"/getAll": getAll,
		"/delete": delete,
	}

	for route, function := range rm {
		http.HandleFunc(route, function)
	}

	log.Fatal(http.ListenAndServe(":8081", nil))
}
