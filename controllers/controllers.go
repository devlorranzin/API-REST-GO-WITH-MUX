package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devlorranzin/apis-golang-tarrafada/database"
	"github.com/devlorranzin/apis-golang-tarrafada/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func ListarEmpresas(w http.ResponseWriter, r *http.Request) {
	var e []models.Empresa
	database.DB.Find(&e)
	json.NewEncoder(w).Encode(e)
}

func ListaEmpresaPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var empresa models.Empresa
	database.DB.First(&empresa, id)
	json.NewEncoder(w).Encode(empresa)
}

func CriaUmaNovaEmpresa(w http.ResponseWriter, r *http.Request) {
	var novaEmpresa models.Empresa
	json.NewDecoder(r.Body).Decode(&novaEmpresa)
	database.DB.Create(&novaEmpresa)
	json.NewEncoder(w).Encode(novaEmpresa)
}

func DeletaUmaEmpresa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var empresa models.Empresa
	database.DB.Delete(&empresa, id)
	json.NewEncoder(w).Encode(empresa)
}

func EditaUmaEmpresa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var empresa models.Empresa
	database.DB.First(&empresa, id)
	json.NewDecoder(r.Body).Decode(&empresa)
	database.DB.Save(&empresa)
	json.NewEncoder(w).Encode(empresa)
}
