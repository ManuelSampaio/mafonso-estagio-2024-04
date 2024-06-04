package produto

type Produto struct {
	id           string
	nome         string
	dataValidade string
}

func New(i, n, d string) *Produto {
	return &Produto{
		id:           i,
		nome:         n,
		dataValidade: d,
	}
}

func (p Produto) GetId() string {
	return p.id
}
