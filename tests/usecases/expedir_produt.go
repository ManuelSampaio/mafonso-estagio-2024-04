package usecases

import (
    "acme_stock_management/internal/domain/entities"
    "testing"
    "time"
)

func TestExpediteProduct(t *testing.T) {
    productRepo := &MockProductRepository{products: make(map[string]*entities.Product)}
    shipmentRepo := &MockShipmentRepository{shipments: make(map[string]*entities.Shipment)}
    usecase := NewExpediteProductUsecase(productRepo, shipmentRepo)

    t.Run("should successfully expedite a product", func(t *testing.T) {
        product := &entities.Product{
            ID:            "1",
            Name:          "Milk",
            Quantity:      10,
            ExpirationDate: time.Now().AddDate(0, 1, 0),
        }
        _ = productRepo.Save(product)

        shipment, err := usecase.ExpediteProduct("1", 5)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        if shipment == nil || shipment.ProductID != "1" || shipment.Quantity != 5 {
            t.Errorf("expected valid shipment, got %v", shipment)
        }

        remainingProduct, _ := productRepo.FindByID("1")
        if remainingProduct.Quantity != 5 {
            t.Errorf("expected remaining quantity to be 5, got %d", remainingProduct.Quantity)
        }
    })

    t.Run("should return error if product quantity is insufficient", func(t *testing.T) {
        product := &entities.Product{
            ID:            "2",
            Name:          "Milk",
            Quantity:      3,
            ExpirationDate: time.Now().AddDate(0, 1, 0),
        }
        _ = productRepo.Save(product)

        _, err := usecase.ExpediteProduct("2", 5)
        if err == nil {
            t.Errorf("expected an error, got none")
        }
    })
}
