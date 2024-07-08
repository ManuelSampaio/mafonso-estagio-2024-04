package application

func AdicionarSubsidiosAoSalario(salario float64, subsidios []float64) float64 {
    for _, subsidio := range subsidios {
        salario += subsidio
    }
    return salario
}

func AplicarDescontoINSS(salario float64, percentualINSS float64) float64 {
    return salario * (1 - percentualINSS/100)
}

func CalcularIRTGrupoA(salarioBase float64) float64 {
    var irt float64

    if salarioBase <= 5000 {
        irt = 0
    } else if salarioBase > 5000 && salarioBase <= 10000 {
        irt = salarioBase * 0.15
    } else {
        irt = salarioBase * 0.25
    }

    return irt
}