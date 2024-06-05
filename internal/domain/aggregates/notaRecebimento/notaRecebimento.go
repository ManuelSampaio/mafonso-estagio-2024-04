package notarecebimento

type NotaRecebimento struct {
	id         string
	quantidade int
	validade   string
}

func New(i string, q int, v string) NotaRecebimento {
	return NotaRecebimento{id: i, quantidade: q, validade: v}
}

func (n NotaRecebimento) Id() string {
	return n.id
}

func (n NotaRecebimento) Quantidade() int {
	return n.quantidade
}

func (n NotaRecebimento) Validade() string {
	return n.validade
}
