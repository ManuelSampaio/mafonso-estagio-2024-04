package expedirMercadoria

import (
	"ACMELDA/internal/domain/aggregates/guiaRemessa"
	"ACMELDA/internal/domain/entities/lote"
	"ACMELDA/internal/domain/repository/guiaRemessaRepository"
)

type ExpedirMercadoria struct {
	lote             lote.Lote
	guiaRepositorio guiaRemessaRepository.GuiaRemessaRepository
}

func (e ExpedirMercadoria) Executa(g *guiaRemessa.GuiaRemessa) bool {
	lotes := e.lote.EncontraProdutoEmUmOuMaisLotes(g.ProdutoID())
	retirou := e.lote.RetirarProdutoNoLote(lotes, g.Quantidade())
	guia := e.guiaRepositorio.New()

	if retirou != -1 {
		guia.CriarGuia(g)
		if existe, _ := guia.RecuperarGuia(g.Id()); existe.Id() != "" {
			return true
		}
	}

	return false
}

/*func (e ExpedirMercadoria) Executa(gr guiaRemessaRepository.GuiaRemessaRepository, g *guiaRemessa.GuiaRemessa) bool {
	lotes := e.lote.EncontraProdutoEmUmOuMaisLotes(g.ProdutoID())
	retirou := e.lote.RetirarProdutoNoLote(lotes, g.Quantidade())
	repositorio := gr.New()

	if retirou != -1 {
		repositorio.CriarGuia(g)

		if existe, _ := repositorio.RecuperarGuia(g.Id()); existe.Id() != "" {
			return true
		}
	}

	return false
}*/
