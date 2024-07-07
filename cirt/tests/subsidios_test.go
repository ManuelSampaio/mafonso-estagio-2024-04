package tests

import (
    "testing"
    "github.com/ManuelSampaio/mafonso-estagio-2024-04/application"
)

func TestAdicionarSubsidios(t *testing.T) {
    t.Run("Teste de adição de subsídios", func(t *testing.T) {
        // Arrange
        salario := 5000.0
        subsidios := []float64{100.0, 50.0}

        // Act
        salarioComSubsidios := application.AdicionarSubsídios(salario, subsidios)

        // Assert
        esperado := 5150.0
        if salarioComSubsidios != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioComSubsidios)
        }
    })
}
