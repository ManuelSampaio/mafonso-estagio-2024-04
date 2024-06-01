package entities
package entidades

import "time"

type Produto struct {
    ID             string
    Nome           string
    Quantidade     int
    DataValidade   time.Time
}
