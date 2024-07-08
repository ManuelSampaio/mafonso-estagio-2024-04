package domain

type Salario struct {
	Base            float64
	Faltas          int
	Subsidios       []float64
	PercentualINSS  float64
	GrupoTributacao int
}
