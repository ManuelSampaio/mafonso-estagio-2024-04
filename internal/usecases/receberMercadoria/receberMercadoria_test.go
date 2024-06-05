package recebermercadoria_test

import (
	notarecebimento "ACMELDA/internal/domain/aggregates/notaRecebimento"
	"testing"
)

func TestReceberMercadoria(t *testing.T) {
	t.Run("se nota de recebimento foi criada", func(t *testing.T) {
		//arrange

		// act
		n := notarecebimento.New("", 0, "")

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
		n := notarecebimento.New(id, Quantidade, validade)

		//assert
		if n.Validade() == "" {
			t.Fail()
		}
	})
}
