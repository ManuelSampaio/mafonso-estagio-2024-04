package application

func CalcularSalarioComFaltas(salarioBase float64, faltas int, descontoPorFalta float64) float64 {
    return salarioBase - (float64(faltas) * descontoPorFalta)
}
