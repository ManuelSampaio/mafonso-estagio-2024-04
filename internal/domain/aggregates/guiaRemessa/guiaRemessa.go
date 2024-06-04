package guiaRemessa

import (
	"time"
)

type GuiaRemessa struct {
	id            string
	pedidoID      string
	quantidade    int
	dataExpedicao time.Time
}

func NewGuiaRemessa(id, pedidoID string, quantidade int) *GuiaRemessa {
	return &GuiaRemessa{
		id:            id,
		pedidoID:      pedidoID,
		quantidade:    quantidade,
		dataExpedicao: time.Now(),
	}
}

func (g *GuiaRemessa) GetID() string {
	return g.id
}

func (g *GuiaRemessa) GetPedidoID() string {
	return g.pedidoID
}

func (g *GuiaRemessa) GetDataExpedicao() time.Time {
	return g.dataExpedicao
}

func (g *GuiaRemessa) GetQuantidade() int {
	return g.quantidade
}
