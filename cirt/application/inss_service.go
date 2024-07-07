package application

func AplicarINSS(salario float64, percentualINSS float64) float64 {
    return salario * (1 - percentualINSS/100)
}
