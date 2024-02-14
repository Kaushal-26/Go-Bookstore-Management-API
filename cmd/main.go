package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kaushal_26/go-bookstore/pkg/routes"
	"github.com/Kaushal_26/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Your .env file where all the varibles are defined
	current_working_directory, err := os.Getwd()
	utils.CheckError(err)
	godotenv.Load(current_working_directory + "..\\.env")

	// PORT Where you want the API to Start defined in .env
	PORT := os.Getenv("PORT")

	// mux router
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)

	fmt.Println("Starting at PORT", PORT, "\n")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, router))
}
