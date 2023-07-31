package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetAndServeRouter sets handler method and starts the application to listen on provided port number
func SetAndServeRouter(config Config) error {
	router := mux.NewRouter() // router for advanced routing capabilities
	handler := newHandler(config)
	router.HandleFunc("/fetchData", handler.FetchVirtualDeviceData).Methods("GET") // Define routes
	server := http.Server{
		Addr:    config.Port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
