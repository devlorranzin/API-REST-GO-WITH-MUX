package main

import (
	"fmt"
	"net/http"

	"github.com/devlorranzin/apis-golang-tarrafada/database"
	"github.com/devlorranzin/apis-golang-tarrafada/routes"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func main() {

	database.Connectbd()

	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
