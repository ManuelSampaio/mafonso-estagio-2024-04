package notarecebimento

type NotaRecebimento struct {
	Id         string
	Quantidade int
	validade   string
}

func New(i string, q int, v string) NotaRecebimento {
	return NotaRecebimento{Id: i, Quantidade: q, validade: v}
}
