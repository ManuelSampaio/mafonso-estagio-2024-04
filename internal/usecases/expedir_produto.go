package usecase

import (
    "errors"

    "ACME,LDA/internal/domain"
)

type Usecase struct {
    repository domain.Repository
    uc         domain.Usecase
}

func NewUsecase(repo domain.Repository, uc domain.Usecase) *Usecase {
    return &Usecase{
        repository: repo,
        uc:         uc,
    }
}

func (uc *Usecase) ExpedirProduto(produtoID string) error {
    produtos, err := uc.repository.GetProdutosParaExpedir()
    if err != nil {
        return err
    }

    for _, produto := range produtos {
        if produto.ID == produtoID {
            guiaRemessa := &domain.GuiaRemessa{
                ID:         produto.ID,
                ProdutoID:  produto.ID,
                ClienteID:  produto.ClienteID,
                DataEnvio:  "DataEnvio",
                Endereco:   produto.Endereco,
                Produto:    produto.Produto,
                Quantidade: produto.Quantidade,
            }

            err = uc.repository.SaveGuiaRemessa(guiaRemessa)
            if err != nil {
                return err
            }

            return nil
        }
    }

    return errors.New("Produto não encontrado")
}
