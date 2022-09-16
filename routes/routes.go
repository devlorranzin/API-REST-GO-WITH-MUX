package routes

import (
	"log"
	"net/http"

	"github.com/devlorranzin/apis-golang-tarrafada/controllers"
	"github.com/devlorranzin/apis-golang-tarrafada/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/empresas", controllers.ListarEmpresas).Methods("Get")
	r.HandleFunc("/empresas/{id}", controllers.ListaEmpresaPorId).Methods("Get")
	r.HandleFunc("/empresas", controllers.CriaUmaNovaEmpresa).Methods("Post")
	r.HandleFunc("/empresas/{id}", controllers.DeletaUmaEmpresa).Methods("Delete")
	r.HandleFunc("/empresas/{id}", controllers.EditaUmaEmpresa).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
