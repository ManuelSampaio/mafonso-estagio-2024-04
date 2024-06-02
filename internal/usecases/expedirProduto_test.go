package expedirProduto_test

import (
		"testing"
		"fmt"
		produto "ACMELDA/internal/domain/entities/produto"
		guia  "ACMELDA/internal/domain/aggregates/guiaRemessa"
		repositorio "ACMELDA/internal/domain/repository/guiaRemessaRepository"
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

	t.Run("se repositorio guia de remessa foi criada", func(t *testing.T){
		// arrange
		p:= produto.Produto{}
		g:= guia.GuiaRemessa{}
		r := repositorio.New()
		p.New("001", "Pao", "2024-06-12")
		g.New("g001", p, 100)
		// act
		gr:= r. CriarRepositorioGuia(&g)
		// assert
		if gr.Id() == "" {
			t.Fail()
		}

	} )

	t.Run("recuperar um agregado de guia remessa", func(t *testing.T){
			// arrange
			p:= produto.Produto{}
			g:= guia.GuiaRemessa{}
			r := repositorio.New()
			p.New("001", "Pao", "2024-06-12")
			g.New("g001", p, 100)
			r. CriarRepositorioGuia(&g)

			// act 
			gr:= r.RecuperarRepositorioGuia(1)

			// assert
			if gr != nil {
				fmt.Println(gr)
				t.Fail()
			}
	} )
}

