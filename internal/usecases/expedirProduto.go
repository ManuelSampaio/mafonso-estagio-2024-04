package expedirProduto

import (
    "ACMELDA/internal/domain/aggregates/guiaRemessa"
    "ACMELDA/internal/domain/repository/guiaRemessaRepository"
)

type ExpedirProdutoInput struct {
    ID        string
    PedidoID  string
    Quantidade int
}

type ExpedirProdutoOutput struct {
    GuiaID      string
    PedidoID    string
    Quantidade  int
    DataExpedicao string
}

type ExpedirProdutoUseCase struct {
    repositorio *guiaRemessaRepository.GuiaRemessaRepository
}

func NewExpedirProdutoUseCase(repo *guiaRemessaRepository.GuiaRemessaRepository) *ExpedirProdutoUseCase {
    return &ExpedirProdutoUseCase{
        repositorio: repo,
    }
}

func (uc *ExpedirProdutoUseCase) Execute(input ExpedirProdutoInput) (*ExpedirProdutoOutput, error) {
    guia := guiaRemessa.NewGuiaRemessa(input.ID, input.PedidoID, input.Quantidade)
    uc.repositorio.CriarGuia(guia)

    return &ExpedirProdutoOutput{
        GuiaID:       guia.GetID(),
        PedidoID:     guia.GetPedidoID(),
        Quantidade:   guia.GetQuantidade(),
        DataExpedicao: guia.GetDataExpedicao().String(),
    }, nil
}


