package repository

import "github.com/ManuelSampaio/mafonso-estagio-2024-04/domain"

type InMemSalarioRepository struct {
	salarios []domain.Salario
}

func NewInMemSalarioRepository() *InMemSalarioRepository {
	return &InMemSalarioRepository{
		salarios: make([]domain.Salario, 0),
	}
}

func (r *InMemSalarioRepository) Create(salario domain.Salario) {
	r.salarios = append(r.salarios, salario)
}

func (r *InMemSalarioRepository) GetAll() []domain.Salario {
	return r.salarios
}
