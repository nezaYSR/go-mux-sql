package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nezaYSR/go-mux-sql/pkg/controllers"
)

var RegisterScrollStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/yippie/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello this is works!")
	}).Methods("GET")
	router.HandleFunc("/scroll/", controllers.CreateScroll).Methods("POST")
	router.HandleFunc("/scroll/", controllers.GetScroll).Methods("GET")
	router.HandleFunc("/scroll/{scrollId}", controllers.GetScrollById).Methods("GET")
	router.HandleFunc("/scroll/{scrollId}", controllers.UpdateScroll).Methods("PUT")
	router.HandleFunc("/scroll/{scrollId}", controllers.DeleteScroll).Methods("DELETE")
}
