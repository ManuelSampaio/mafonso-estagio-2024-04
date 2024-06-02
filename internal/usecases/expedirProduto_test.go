package expedirProduto_test

import (
		"testing"
		produto "ACMELDA/internal/domain/entities/produto"
		guia  "ACMELDA/internal/domain/aggregates/guiaRemessa"
)

func testExpedirMercadoria(t *testing.T){

	t.Run("se guia de remessa foi criada", func(t *testing.T){
		// arrange
		p:= produto.Produto{}
		
		p.New("001", "Pao", "2024-06-12")

		// act
		g:= guia.GuiaRemessa{}

		g.New("g001", p, 100)
		// assert

		if g.Id() == "" {
			t.Fail()
		}

	} )
}

