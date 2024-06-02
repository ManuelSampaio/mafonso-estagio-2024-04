package guiaRemessarepository

import (
      guia  "ACMELDA/internal/domain/aggregates/guiaRemessa"
)

type GuiaRemessaRepository interface {
    CriaGuiaRemessa(guia *guia.GuiaRemessa) (*guia.GuiaRemessa, error)
}