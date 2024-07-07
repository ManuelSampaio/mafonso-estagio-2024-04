package tests

import (
    "testing"
    "github.com/ManuelSampaio/mafonso-estagio-2024-04/application"
)

func TestAplicarINSS(t *testing.T) {
    t.Run("Teste de aplicação de INSS", func(t *testing.T) {
        // Arrange
        salario := 5000.0
        percentualINSS := 8.0

        // Act
        salarioComINSS := application.AplicarINSS(salario, percentualINSS)

        // Assert
        esperado := 4600.0
        if salarioComINSS != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioComINSS)
        }
    })
}
