package guiaremessa

import (
	produto "Acme/domain/entities/produto"
)

type GuiaRemessa struct {
	Id         string
	Produtos   produto.Produto
	Quantidade int
}
