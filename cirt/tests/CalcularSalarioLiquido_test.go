package tests

import (
	"testing"
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/application"
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/domain"
)

func TestCalcularSalarioLiquido(t *testing.T) {
	t.Run("Teste de cálculo de salário líquido com IRT - Grupo A", func(t *testing.T) {
		// Arrange
		salario := domain.Salario{
			Base:            5000.0,
			Subsidios:       []float64{100.0, 50.0},
			PercentualINSS:  8.0,
			GrupoTributacao: "A",
		}

		// Act
		salarioLiquido := application.CalcularSalarioLiquido(salario)

		// Assert
		esperado := 4550.0
		if salarioLiquido != esperado {
			t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioLiquido)
		}
	})

	t.Run("Teste de cálculo de salário líquido sem IRT", func(t *testing.T) {
		// Arrange
		salario := domain.Salario{
			Base:            5000.0,
			Subsidios:       []float64{100.0, 50.0},
			PercentualINSS:  8.0,
			GrupoTributacao: "B",
		}

		// Act
		salarioLiquido := application.CalcularSalarioLiquido(salario)

		// Assert
		esperado := 4860.0
		if salarioLiquido != esperado {
			t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioLiquido)
		}
	})
}

