package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CezarGarrido/FarmVue/ApiFarm/driver"
	areaHandler "github.com/CezarGarrido/FarmVue/ApiFarm/handlers"
	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
)

func main() {
	dbUser := "postgres"
	dbPass := "C102030g"
	dbName := "api_farm_repo"
	dbHost := "localhost"
	dbPort := "5432"
	connection, err := driver.ConexaoPostgres(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	aHandler := areaHandler.NewAreaHandler(connection)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/area", aHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/area/{id:[0-9]+}", aHandler.Delete).Methods("DELETE")
	r.HandleFunc("/api/v1/areas/proprietario/{id:[0-9]+}", aHandler.GetAllByProprietarioID).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD","DELETE","OPTIONS"})
	fmt.Println("Servidor startado na porta :4000")

	http.ListenAndServe(":4000", handlers.CORS(headersOk, methodsOk, originsOk)(r))
}
