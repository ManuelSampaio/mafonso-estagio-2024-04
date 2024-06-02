package repository

import (
	"errors"

	"github.com/seu-user/seu-projeto/internal/domain"
)

type Repository struct {
	database domain.Database
}

func NewRepository(db domain.Database) *Repository {
	return &Repository{
		database: db,
	}
}

func (repo *Repository) GetProdutosParaExpedir() ([]*domain.Produto, error) {
	// Implemente a lógica para buscar os produtos para expedir
	// ...
	// Exemplo: return repo.database.GetProdutosParaExpedir()
	// ...
	// Caso não encontre produtos, retorne um erro personalizado
	return nil, errors.New("Não há produtos para expedir")
}

func (repo *Repository) SaveGuiaRemessa(guiaRemessa *domain.GuiaRemessa) error {
	// Implemente a lógica para salvar a guia de remessa
	// ...
	// Exemplo: return repo.database.SaveGuiaRemessa(guiaRemessa)
	// ...
	// Caso ocorra algum erro ao salvar a guia de remessa, retorne o erro
	return errors.New("Erro ao salvar guia de remessa")
}