package application

import (
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/domain"
)


func CalcularSalarioLiquido(salario domain.Salario) float64 {
	
	salarioCalculado := AdicionarSubsidios(salario.Base, salario.Subsidios)

	salarioCalculado = AplicarINSS(salarioCalculado, salario.PercentualINSS)

	if salario.GrupoTributacao == "A" {
		salarioCalculado = CalcularIRTGrupoA(salarioCalculado)
	}

	return salarioCalculado
}

func AdicionarSubsidios(salario float64, subsidios []float64) float64 {
	for _, subsidio := range subsidios {
		salario += subsidio
	}
	return salario
}


func AplicarINSS(salario float64, percentualINSS float64) float64 {
	return salario * (1 - percentualINSS/100)
}


func CalcularIRTGrupoA(salario float64) float64 {
	
	return salario
}
