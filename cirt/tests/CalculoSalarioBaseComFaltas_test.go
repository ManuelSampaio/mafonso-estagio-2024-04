package tests

import (
    "testing"

    "cirt/application"

)

func TestCalcularSalarioBaseComFaltas(t *testing.T) {

    t.Run("Teste sem faltas", func(t *testing.T) {
    
        salarioBase := 5000.0
        faltas := 0

        // Act
        salarioCalculado := application.CalcularSalarioComFaltas(salarioBase, faltas)

        // Assertd
        esperado := 5000.0
        if salarioCalculado != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioCalculado)
        }
    })


    t.Run("Teste com faltas", func(t *testing.T) {
        // Arrange
        salarioBase := 5000.0
        faltas := 2

        // Act
        salarioCalculado := application.CalcularSalarioComFaltas(salarioBase, faltas)

        // Assert
        esperado := 4900.0
        if salarioCalculado != esperado {
            t.Errorf("Resultado esperado %f, mas obtido %f", esperado, salarioCalculado)
        }
    })

    
}
