package recebermercadoria_test

import (
	notarecebimento "ACMELDA/internal/domain/aggregates/notaRecebimento"
	notarecebimentorepository "ACMELDA/internal/domain/repository/notaRecebimentoRepository"
	"testing"
)

func TestReceberMercadoria(t *testing.T) {
	t.Run("se nota de recebimento foi criada", func(t *testing.T) {
		//arrange

		// act
		n := notarecebimento.New("", "", 0, "")

		//assert
		if n.Quantidade() != 0 {
			t.Fail()
		}
	})

	t.Run("se há dados na nota de recebimento", func(t *testing.T) {
		//arrange
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"
		// act
		n := notarecebimento.New(id, "", Quantidade, validade)

		//assert
		if n.Validade() == "" {
			t.Fail()
		}
	})

	t.Run("se há produto na nota de recebimento", func(t *testing.T) {
		//arrange
		produtoID := "013"
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"
		// act
		n := notarecebimento.New(id, produtoID, Quantidade, validade)

		//assert
		if n.ProdutoID() == "" {
			t.Fail()
		}
	})

	t.Run("recuperar agregado no repositorio da nota de recebimento", func(t *testing.T) {
		//arrange
		produtoID := "013"
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"

		n := notarecebimento.New(id, produtoID, Quantidade, validade)
		repo := notarecebimentorepository.New()
		repo.CriarNotaRecebimento(&n)

		// act
		rp, err := repo.RecuperarNotaRecebimento(id)

		//assert
		if err != nil || rp == nil {

			t.Fail()
		}

	})

	t.Run("se criou mais de um agregado no repositorio da nota de recebimento", func(t *testing.T) {
		//arrange
		var notasRecebimento []notarecebimento.NotaRecebimento
		repo := notarecebimentorepository.New()

		notasRecebimento = append(notasRecebimento, notarecebimento.New("n001", "002", 1672, "2025-02-11"))
		notasRecebimento = append(notasRecebimento, notarecebimento.New("n001", "005", 1432, "2025-03-21"))
		notasRecebimento = append(notasRecebimento, notarecebimento.New("n001", "022", 112, "2025-04-01"))
		notasRecebimento = append(notasRecebimento, notarecebimento.New("n001", "202", 120, "2025-05-22"))

		for _, nota := range notasRecebimento {
			repo.CriarNotaRecebimento(&nota)
		}

		// act
		rp := repo.RecuperarTodasNotasRecebimento()

		//assert
		if len(rp) == 0 {
			t.Fail()
		}

	})

	t.Run("nota de recebimento sem dado", func(t *testing.T) {
		//arrange

		//act
		n := notarecebimento.New("", "", 0, "")

		// assert
		if n.Quantidade() > 0 {
			//fmt.Println("NOTA: ", n)
			t.Fail()
		}

	})
	// depois de recebido o produto é armazenado (LOTES: Pratileira-Corredor-Armario)
	// ao receber : terei o produto, depois irei colocar em uma nota de recebimento,
	//depois recuperar essa nota (depois de criada, o produto é armazenado em lote)

	// intergarar algumas funcionalidade
}
