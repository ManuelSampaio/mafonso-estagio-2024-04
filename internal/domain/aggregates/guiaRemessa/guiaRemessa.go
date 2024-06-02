package guiaremessa

import (
	produto "ACMELDA/internal/domain/entities/produto"
)

type GuiaRemessa struct {
	id         string
	produtos   produto.Produto
	quantidade int
}

func (g *GuiaRemessa) New(i string, p produto.Produto, q int) *GuiaRemessa{
	g.id =i
	g.produtos = p
	g.quantidade = q

	return g
}

func (g GuiaRemessa) Id() string{
	return g.id
}
