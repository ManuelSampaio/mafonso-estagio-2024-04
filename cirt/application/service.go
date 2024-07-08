package application

import (
	"fmt"
	"github.com/ManuelSampaio/mafonso-estagio-2024-04/domain"
)

const (
	GrupoA = iota + 1
	GrupoB
)

func CalcularSalarioLiquido(salario domain.Salario) (float64, error) {
    salarioCalculado := CalcularSalarioComFaltas(salario.Base, salario.Faltas, 50.0)
    salarioCalculado = AdicionarSubsidios(salarioCalculado, salario.Subsidios)
    salarioCalculado = AplicarINSS(salarioCalculado, salario.PercentualINSS)

    switch salario.GrupoTributacao {
    case GrupoA:
        salarioCalculado = CalcularIRTGrupoA(salarioCalculado)
    case GrupoB:
        salarioCalculado = CalcularIRTGrupoB(salarioCalculado)
    default:
        return 0, fmt.Errorf("grupo de tributação inválido: %d", salario.GrupoTributacao)
    }

    return salarioCalculado, nil
}
