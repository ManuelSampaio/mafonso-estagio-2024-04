package lote

import (
	produto "ACMELDA/internal/domain/entities/produto"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Lote struct {
	Id           string
	Produto      produto.Produto
	Quantidade   int
	DataValidade string
}

func NewLote() *Lote {
	return &Lote{
		Id:           "",
		Produto:      produto.Produto{},
		Quantidade:   0,
		DataValidade: "",
	}
}

func (l Lote) EncontraProdutoEmUmOuMaisLotes(idproduto string) []Lote {
	Dadoslotes := l.lerDadosArquivo()
	lotes := l.converteEmStruct(Dadoslotes)
	var loteaulixiar []Lote

	for _, lote := range lotes {
		auxiliar1, _ := strconv.Atoi(lote.Produto.GetId())
		auxiliar2, _ := strconv.Atoi(idproduto)

		if auxiliar1 == auxiliar2 && DataValidadeProxima(lote.DataValidade) {
			fmt.Println(" ENCONTREI: ", lote)
			loteaulixiar = append(loteaulixiar, lote)
		}

	}
	return loteaulixiar

}

func (l Lote) RetirarProdutoNoLote(lotes []Lote, quantidade int) int {

	unidadeSubstituta := quantidade

	for _, lote := range lotes {
		fmt.Println(" O LOTE QUE DSCONCTOU Principio: ", lote)
		if lote.Quantidade >= unidadeSubstituta && unidadeSubstituta > 0 {
			lote.Quantidade = lote.Quantidade - unidadeSubstituta
			fmt.Println(" SERÀ???")
			return lote.Quantidade

		} else if lote.Quantidade <= unidadeSubstituta && unidadeSubstituta > 0 {
			fmt.Println(" OOIIIII LOte: ", lote.Id)
			unidadeSubstituta = unidadeSubstituta - lote.Quantidade
			lote.Quantidade -= lote.Quantidade
			fmt.Println(" FIM LOte: ", lote)
		}
	}

	return -1
}

func (l Lote) lerDadosArquivo() string {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	defer func() {
		err = os.Chdir(dir)
		if err != nil {
			fmt.Println("Erro ao voltar ao diretório de trabalho:", err)
		}
	}()

	nomeArquivo := "lote.txt"
	novoDir := filepath.Join(dir, "..", "domain", "entities", "lote")
	err = os.Chdir(novoDir)

	if err != nil {
		fmt.Println("Erro ao mudar de diretório2:", err)
		return ""
	}

	dados, err := ioutil.ReadFile(novoDir + "/" + nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao ler os dados:", err)
		return ""
	}
	novoDir = ""

	return string(dados)
}

func (l Lote) converteEmStruct(dados string) []Lote {

	blocos := strings.Split(dados, "\n\n")
	loteMap := make(map[int]Lote)
	idLote := 1

	for _, bloco := range blocos {
		if bloco != "" {
			linhas := strings.Split(bloco, "\n")

			lote := Lote{}

			for _, linha := range linhas {
				campos := strings.SplitN(linha, ": ", 2)
				if len(campos) == 2 {
					switch campos[0] {
					case "Id":
						lote.Id = campos[1]
					case "Produto":
						produtoCampos := strings.Split(campos[1], ", ")
						produtoId, _ := strconv.Atoi(produtoCampos[0][4:])
						produtoNome := produtoCampos[1][6:]
						dataValidadeProduto := produtoCampos[2][14:]
						lote.Produto = *produto.New(strconv.Itoa(produtoId), produtoNome, strings.Replace(dataValidadeProduto, "}", "", 1))

					case "Quantidade":
						lote.Quantidade, _ = strconv.Atoi(strings.Replace(campos[1], ",", "", 1))

					case "DataValidade":
						lote.DataValidade = campos[1]
					}
				}
			}

			if _, ok := loteMap[idLote]; !ok {
				loteMap[idLote] = lote
				idLote += 1
			}
		}
	}

	lotes := []Lote{}
	for _, lote := range loteMap {
		lotes = append(lotes, lote)
	}

	return lotes
}
