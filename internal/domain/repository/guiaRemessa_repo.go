package repositorios

import "acme,lda/internal/domain/entities"

type GuiaRemessaRepositorio interface {
    Salvar(guiaRemessa *entidades.GuiaRemessa) error
    BuscarPorID(id string) (*entidades.GuiaRemessa, error)
}
