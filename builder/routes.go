package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func drawRoutes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", Index)
	fmt.Println("wiring up routes")

	return router
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("something fishy here")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
