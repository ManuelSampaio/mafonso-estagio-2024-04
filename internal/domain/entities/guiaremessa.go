package entities

import "time"

type GuiaRemessa struct {
    ID            string
    ProdutoID     string
    Quantidade    int
    DataExpedicao time.Time
}
