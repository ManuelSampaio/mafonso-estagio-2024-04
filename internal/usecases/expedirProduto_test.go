package expedirProduto_test

import (
	"ACMELDA/internal/domain/aggregates/guiaRemessa"
	"ACMELDA/internal/domain/entities/lote"
	"ACMELDA/internal/domain/repository/guiaRemessaRepository"
	"testing"
)

func TestExpedirMercadoria(t *testing.T) {

	t.Run("se guia de remessa foi criada", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 100

		// act
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)

		// assert
		if g.GetID() == "" {
			t.Fail()
		}
	})

	t.Run("se repositorio guia de remessa foi criada", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 100
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		r := guiaRemessaRepository.New()

		// act
		r.CriarGuia(g)

		// assert
		if g.GetID() == "" {
			t.Fail()
		}
	})

	t.Run("recuperar um agregado de guia remessa", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 100
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		r := guiaRemessaRepository.New()
		r.CriarGuia(g)

		// act
		gr, err := r.RecuperarGuia("g001")

		// assert
		if err != nil || gr == nil {
			t.Fatalf("Erro ao recuperar guia de remessa: %v", err)
		}
	})

	t.Run("se o produto for encontrado no lote", func(t *testing.T) {
		// arrange
		pedidoID := "003"
		quantidade := 100
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.NewLote()

		//act
		p := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())

		// assert
		if len(p) == 0 {
			t.Fail()
		}

	})

	t.Run("retorna produtos em lotes com data validade mais próxima", func(t *testing.T) {
		// arrange
		pedidoID := "003"
		quantidade := 30
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.NewLote()

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())

		// assert
		if len(lotes) == 0 {
			t.Fail()
		}

	})

	t.Run("buscar um produto que não existe em lotes", func(t *testing.T) {
		// arrange
		pedidoID := "103"
		quantidade := 0
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.NewLote()

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())

		// assert
		if len(lotes) != 0 {
			t.Fail()
		}

	})
}
