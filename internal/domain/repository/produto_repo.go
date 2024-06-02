package repositories

import (
    "domain/entities"
)

type ProdutoRepository interface {
    GetProdutoByID(string) (*entities.Produto, error)
    AtualizarInventario(string, int) error
}
