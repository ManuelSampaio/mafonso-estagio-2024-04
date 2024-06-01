package expedirproduto_test

import (
	guia "Acme/domain/aggregates/guiaRemessa"
	lote "Acme/domain/entities/lote"
	produto "Acme/domain/entities/produto"
	"fmt"
	"testing"
)

func TestExpedirMercadoria(t *testing.T) {
	t.Run("teste expedir mercadoria", func(t *testing.T) {
		// arrange
		guia := guia.GuiaRemessa{
			Id:         "g001",
			Produtos:   produto.Produto{Id: "p001", Nome: "pao"},
			Quantidade: 29,
		}
		// buscar o produto no lote

		lote := lote.Lote{}
		lote.RetornaTodosLotes()
		fmt.Println(guia)

		if 1 == 1 {
			t.Fail()
		}

	})
}
