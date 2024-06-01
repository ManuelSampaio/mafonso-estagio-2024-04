package main

import (
    "acme_stock_management/internal/interfaces/api/http"
    "acme_stock_management/internal/usecases"
    "net/http"
)

func main() {
    productRepo := &usecases.MockProductRepository{products: make(map[string]*entities.Product)}
    shipmentRepo := &usecases.MockShipmentRepository{shipments: make(map[string]*entities.Shipment)}

    expediteProductUsecase := usecases.NewExpediteProductUsecase(productRepo, shipmentRepo)
    productHandler := http.NewProductHandler(expediteProductUsecase)

    http.HandleFunc("/expedite", productHandler.HandleExpediteProduct)
    http.ListenAndServe(":8080", nil)
}
