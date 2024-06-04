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

	t.Run("Não retornar produtos em lotes com data validade já passada", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 30
		g := guiaRemessa.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.NewLote()

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())

		// assert
		if len(lotes) > 0 {
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
		if len(lotes) > 0 {
			t.Fail()
		}

	})

	t.Run("encontrar produto em lotes sem dados da guia de remessa", func(t *testing.T) {
		// arrange
		g := guiaRemessa.NewGuiaRemessa("", "", 0)
		l := lote.NewLote()

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())

		// assert
		if len(lotes) > 0 {
			t.Fail()
		}

	})

	t.Run("retirar uma determinada quantidade de um produto em lote", func(t *testing.T) {
		// arrange
		produtoID := "003"
		quantidade := 1200

		g := guiaRemessa.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.NewLote()
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, g.GetQuantidade())

		// assert
		if lote == -1 {
			t.Fail()
		}

	})

	t.Run("retirar quantidade acima de um produto mais do que existe no lote", func(t *testing.T) {
		// arrange
		produtoID := "003"
		quantidade := 2900

		g := guiaRemessa.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.NewLote()
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, g.GetQuantidade())

		// assert
		if lote != -1 {
			t.Fail()
		}

	})

	t.Run("retirar um produto que não existe no lote", func(t *testing.T) {
		// arrange
		produtoID := "103"
		quantidade := 900

		g := guiaRemessa.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.NewLote()
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, g.GetQuantidade())

		// assert
		if lote != -1 {
			t.Fail()
		}

	})

	t.Run("retirar um produto sem dados da guia de remessa", func(t *testing.T) {
		// arrange
		quantidade := 199
		g := guiaRemessa.NewGuiaRemessa("", "", 0)
		l := lote.NewLote()
		lotes := l.EncontraProdutoEmUmOuMaisLotes(g.GetPedidoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, quantidade)

		// assert
		if lote != -1 {
			t.Fail()
		}

	})
}
