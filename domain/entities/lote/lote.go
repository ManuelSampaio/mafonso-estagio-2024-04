package lote

import (
	produto "Acme/domain/entities/produto"
	"fmt"
	"io/ioutil"
	"strings"
)

type Lote struct {
	Id           string
	Produto      produto.Produto
	Quantidade   int
	DataValidade string
}

func (l Lote) RetornaTodosLotes() {
	lotes, _ := lerDadosArquivo()
	fmt.Println(lotes)

}

func lerDadosArquivo() ([]string, error) {
	nomeArquivo  := "lote.txt"
	// Lê todo o conteúdo do arquivo
	dados, err := ioutil.ReadFile(nomeArquivo)
	if err != nil {
		return nil, err
	}

	// Divide os dados em linhas
	linhas := strings.Split(string(dados), "\n")
	return linhas, nil
}


