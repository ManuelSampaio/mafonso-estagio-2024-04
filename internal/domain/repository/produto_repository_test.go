package repository_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestSaveGuiaRemessa(t *testing.T) {
	// Repositório sem erros
	repo := repository.NewRepository(&MockDatabase{})
	err := repo.SaveGuiaRemessa(&domain.GuiaRemessa{ID: "ID1"})
	assert.Nil(t, err)

	// Repositório com erro: problemas ao salvar a guia de remessa
	repo = repository.NewRepository(&MockDatabase{err: errors.New("Erro ao salvar guia de remessa")})
	err = repo.SaveGuiaRemessa(&domain.GuiaRemessa{ID: "ID1"})
	assert.NotNil(t, err)
}