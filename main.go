package main

import (
	"fmt"
	"net/http"
	"receipt-processor/api"
	"receipt-processor/service"
	"receipt-processor/storage"
)

func main() {
	fmt.Println("Server is initializing...")
	getRepo, storeRepo := storage.NewReceiptRepository()
	getPoints, submit := service.NewReceiptProcessorService(getRepo, storeRepo)
	rw := api.DefaultResponseWriter{}
	getHandler := api.NewGetReceiptPointsHandler(getPoints, rw)
	submitHandler := api.NewSubmitReceiptHandler(submit, rw)
	getHandler.RegisterRoutes()
	submitHandler.RegisterRoutes()
	fmt.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
