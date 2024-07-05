package application

import "testing"

func TestCalcularSalarioBase(t *testing.T) {
	service := CalculoSalarioService{}

	
	salarioBase := 3000.0

	resultado := service.CalcularSalarioBase(salarioBase)


	valorEsperado := 3000.0
	if resultado != valorEsperado {
		t.Errorf("Resultado esperado %f Kz, mas obtive %f Kz", valorEsperado, resultado)
	}
}
