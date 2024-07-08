package tests

import (
    "testing"
    "github.com/ManuelSampaio/mafonso-estagio-2024-04/application"
)

func TestAdicionarSubsidiosAoSalario(t *testing.T) {
    t.Run("Teste de adição de subsídios", func(t *testing.T) {
        // Arrange
        salario := 5000.0
        subsidios := []float64{100.0, 50.0}

        // Act
        salarioComSubsidios := application.AdicionarSubsidiosAoSalario(salario, subsidios)

        // Assert
        esperado := 5150.0
        if salarioComSubsidios != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioComSubsidios)
        }
    })
}

func TestAplicarDescontoINSS(t *testing.T) {
    t.Run("Teste de aplicação de INSS", func(t *testing.T) {
        // Arrange
        salario := 5000.0
        percentualINSS := 8.0

        // Act
        salarioComINSS := application.AplicarDescontoINSS(salario, percentualINSS)

        // Assert
        esperado := 4600.0
        if salarioComINSS != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioComINSS)
        }
    })
}

func TestCalcularIRTGrupoA(t *testing.T) {
    t.Run("Teste de cálculo do IRT para Grupo A", func(t *testing.T) {
        // Arrange
        salarioBase := 8000.0

        // Act
        irtCalculado := application.CalcularIRTGrupoA(salarioBase)

        // Assert
        esperado := 1200.0 // 8000 * 0.15
        if irtCalculado != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, irtCalculado)
        }
    })
}
