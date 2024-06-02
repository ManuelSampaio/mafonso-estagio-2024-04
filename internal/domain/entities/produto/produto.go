package produto

type Produto struct {
	id           string
	nome         string
	dataValidade string
}

func (p *Produto) New(i, n, d string) *Produto {
	return &Produto{
		   id: i,
		   nome: n,
		   dataValidade: d,
	}
}