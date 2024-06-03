package expedirProduto_test

import (
    "testing"
    "fmt"
    "ACMELDA/internal/domain/aggregates/guiaRemessa"
    "ACMELDA/internal/domain/repository/guiaRemessaRepository"
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
        fmt.Println(gr)
    })
}
