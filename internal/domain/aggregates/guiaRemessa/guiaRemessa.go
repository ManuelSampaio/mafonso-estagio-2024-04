package guiaRemessa

import (
    "time"
)

type GuiaRemessa struct {
    ID            string
    PedidoID      string
    Quantidade    int
    DataExpedicao time.Time
}

func NewGuiaRemessa(id, pedidoID string, quantidade int) *GuiaRemessa {
    return &GuiaRemessa{
        ID:            id,
        PedidoID:      pedidoID,
        Quantidade:    quantidade,
        DataExpedicao: time.Now(),
    }
}

func (g *GuiaRemessa) GetID() string {
    return g.ID
}

func (g *GuiaRemessa) GetPedidoID() string {
    return g.PedidoID
}

func (g *GuiaRemessa) GetDataExpedicao() time.Time {
    return g.DataExpedicao
}

func (g *GuiaRemessa) GetQuantidade() int {
    return g.Quantidade
}
