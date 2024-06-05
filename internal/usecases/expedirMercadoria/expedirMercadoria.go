package expedirMercadoria

import (
	"ACMELDA/internal/domain/aggregates/guiaRemessa"
	"ACMELDA/internal/domain/entities/lote"
)

type ExpedirMercadoria struct {
	lote lote.Lote
}

func (e ExpedirMercadoria) Executa(g *guiaRemessa.GuiaRemessa) bool {
	lotes := e.lote.EncontraProdutoEmUmOuMaisLotes(g.PedidoID())
	retirou := e.lote.RetirarProdutoNoLote(lotes, g.Quantidade())

	return retirou != -1

}
