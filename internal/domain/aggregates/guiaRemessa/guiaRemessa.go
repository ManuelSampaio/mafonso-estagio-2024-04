package guiaremessa

import (
	produto "ACMELDA/internal/domain/entities/produto"
)

type GuiaRemessa struct {
	Id         string
	Produtos   produto.Produto
	Quantidade int
}
