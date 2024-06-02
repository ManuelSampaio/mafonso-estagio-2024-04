package guiaRemessarepository

import (
      guia  "ACMELDA/internal/domain/aggregates/guiaRemessa"
      "sync"
)

type GuiaRemessaRepository struct {
    guias  map[int] *guia.GuiaRemessa
    nextId int
    mutex sync.Mutex
}

func New() * GuiaRemessaRepository {

    return &GuiaRemessaRepository{
        guias: make(map[int] *guia.GuiaRemessa),
        nextId: 1,
    }
}

func (repo *GuiaRemessaRepository) CriarRepositorioGuia(guia *guia.GuiaRemessa) (*guia.GuiaRemessa){
    repo.mutex.Lock()
	defer repo.mutex.Unlock()

	id := repo.nextId
	repo.nextId++
	repo.guias[id] = guia
	return guia
}

func (repo *GuiaRemessaRepository) RecuperarRepositorioGuia(id int) (*guia.GuiaRemessa){

    repo.mutex.Lock()
	defer repo.mutex.Unlock()

    guia, exists := repo.guias[id]
    
	if !exists {
		return nil
    }
    
	return guia
}