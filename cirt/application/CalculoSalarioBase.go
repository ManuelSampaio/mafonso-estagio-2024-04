package application


func CalcularSalarioComFaltas(salarioBase float64, faltas int) float64 {
    
    salarioReduzido := salarioBase - float64(faltas)*50.0
    return salarioReduzido
}
