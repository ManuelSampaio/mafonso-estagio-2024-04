package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ACME,LDA/internal/domain"
	"ACME,LDA/internal/usecase"
)

// MockRepository é um mock da interface domain.Repository
type MockRepository struct {
	Produtos     []domain.Produto
	GuiaRemessa  *domain.GuiaRemessa
	GetProdError error
	SaveError    error
}

func (m *MockRepository) GetProdutosParaExpedir() ([]domain.Produto, error) {
	if m.GetProdError != nil {
		return nil, m.GetProdError
	}
	return m.Produtos, nil
}

func (m *MockRepository) SaveGuiaRemessa(guia *domain.GuiaRemessa) error {
	if m.SaveError != nil {
		return m.SaveError
	}
	m.GuiaRemessa = guia
	return nil
}

// MockUsecase é um mock da interface domain.Usecase
type MockUsecase struct {
	err error
}

func (m *MockUsecase) SomeUsecaseMethod() error {
	return m.err
}

func TestExpedirProduto(t *testing.T) {
	// Caso de uso sem erros
	repo := &MockRepository{
		Produtos: []domain.Produto{
			{
				ID:         "ID1",
				ClienteID:  "Cliente1",
				Endereco:   "Endereco1",
				Produto:    "Produto1",
				Quantidade: 1,
			},
		},
	}
	uc := usecase.NewUsecase(repo, &MockUsecase{})
	err := uc.ExpedirProduto("ID1")
	assert.Nil(t, err)

	// Caso de uso com erro: não há produtos para expedir
	repo = &MockRepository{
		GetProdError: errors.New("Erro ao buscar produtos para expedir"),
	}
	uc = usecase.NewUsecase(repo, &MockUsecase{})
	err = uc.ExpedirProduto("ID1")
	assert.NotNil(t, err)
	assert.Equal(t, "Erro ao buscar produtos para expedir", err.Error())

	// Caso de uso com erro: produto não encontrado
	repo = &MockRepository{
		Produtos: []domain.Produto{},
	}
	uc = usecase.NewUsecase(repo, &MockUsecase{})
	err = uc.ExpedirProduto("ID1")
	assert.NotNil(t, err)
	assert.Equal(t, "Produto não encontrado", err.Error())
}
