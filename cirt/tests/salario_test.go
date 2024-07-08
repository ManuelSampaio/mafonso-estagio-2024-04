package tests

import (
	"testing"

	"github.com/ManuelSampaio/mafonso-estagio-2024-04/application"
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/domain"
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/repository"
)

func TestCalcularSalarioLiquido(t *testing.T) {
    repo := repository.NewInMemSalarioRepository()
    salario := domain.Salario{
        Base:            5000.0,
        Faltas:          2,
        Subsidios:       []float64{100.0, 50.0},
        PercentualINSS:  8.0,
        GrupoTributacao: application.GrupoA,
    }
    repo.Create(salario)

    salarios := repo.GetAll()
    if len(salarios) != 1 {
        t.Errorf("Esperava um salário na lista, mas obteve %d", len(salarios))
    }

    salarioCalculado, err := application.CalcularSalarioLiquido(salario)
    if err != nil {
        t.Errorf("Erro ao calcular salário líquido: %v", err)
    }

    expected := 4590.0 // Valor esperado do salário líquido após cálculos
    if salarioCalculado != expected {
        t.Errorf("Valor esperado do salário líquido era %f, mas obtido foi %f", expected, salarioCalculado)
    }
}
