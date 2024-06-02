package produto

type Produto struct {
	id           string
	nome         string
	dataValidade string
}

func (p *Produto) New(i, n, d string) *Produto {
	p.id = i
	p.nome = n
	p.dataValidade = d
	return p
}