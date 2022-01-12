package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found...")
	}
}

func main() {
	fmt.Println("Hello world!")

	// Setup config
	setupConfig()

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":5000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
