package guiaRemessaRepository

import (
    "ACMELDA/internal/domain/aggregates/guiaRemessa"
    "errors"
)

type GuiaRemessaRepository struct {
    guias map[string]*guiaRemessa.GuiaRemessa
}

func New() *GuiaRemessaRepository {
    return &GuiaRemessaRepository{guias: make(map[string]*guiaRemessa.GuiaRemessa)}
}

func (r *GuiaRemessaRepository) CriarGuia(g *guiaRemessa.GuiaRemessa) {
    r.guias[g.GetID()] = g
}

func (r *GuiaRemessaRepository) RecuperarGuia(id string) (*guiaRemessa.GuiaRemessa, error) {
    g, existe := r.guias[id]
    if !existe {
        return nil, errors.New("guia de remessa não encontrada")
    }
    return g, nil
}
