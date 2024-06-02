package expedirProduto_test

import (
		"testing"
		produto "ACMELDA/internal/domain/entities/produto"
		guia  "ACMELDA/internal/domain/aggregates/guiaRemessa"
)

func testExpedirMercadoria(t *testing.T){

	t.Run("se guia de remessa foi criada", func(t *testing.T){
		// arrange
		p:= produto.Produto{
				 Id: "001",
				 Nome: " Pao",
				DataValidade: "2024-06-12", 
		}

		
		// act
		g:= guia.GuiaRemessa{
			Id: "g001",         
			Produtos: p,
			Quantidade: 100,
		}
		// assert

		if g.Id == "" {
			t.Fail()
		}

	} )
}

