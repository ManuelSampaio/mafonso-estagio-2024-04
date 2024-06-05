package notarecebimentorepository

import (
	notarecebimento "ACMELDA/internal/domain/aggregates/notaRecebimento"
	"errors"
)

type NotaRecebimentoRepository struct {
	notasRecebimento map[string]*notarecebimento.NotaRecebimento
}

func New() *NotaRecebimentoRepository {
	return &NotaRecebimentoRepository{notasRecebimento: make(map[string]*notarecebimento.NotaRecebimento)}

}

func (n *NotaRecebimentoRepository) CriarNotaRecebimento(nt *notarecebimento.NotaRecebimento) {
	n.notasRecebimento[nt.Id()] = nt
}

func (n *NotaRecebimentoRepository) RecuperarNotaRecebimento(id string) (*notarecebimento.NotaRecebimento, error) {
	nt, existe := n.notasRecebimento[id]
	if !existe {
		return nil, errors.New("nota de recebimento não encontrada")
	}
	return nt, nil
}

func (n *NotaRecebimentoRepository) RecuperarTodasNotasRecebimento() []*notarecebimento.NotaRecebimento {
	var todasNotasDeRecebimento []*notarecebimento.NotaRecebimento
	for _, notas := range n.notasRecebimento {
		todasNotasDeRecebimento = append(todasNotasDeRecebimento, notas)
	}
	return todasNotasDeRecebimento
}
