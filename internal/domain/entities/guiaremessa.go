package entities

import "time"

// Shipment represents a shipment entity.
type guiremessa struct {
    ID        string
    ProdutoID string
    Quantidade  int
    CreatedAt time.Time
}
