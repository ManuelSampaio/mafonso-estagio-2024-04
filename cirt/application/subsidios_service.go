package application

func AdicionarSubsidios(salario float64, subsidios []float64) float64 {
    for _, subsidio := range subsidios {
        salario += subsidio
    }
    return salario
}
