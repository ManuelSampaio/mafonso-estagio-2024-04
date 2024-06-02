package lote

import (
	produto "ACMELDA/internal/domain/entities/produto"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"path/filepath"
)

type Lote struct {
	Id           string
	Produto      produto.Produto
	Quantidade   int
	DataValidade string
}

func (l Lote) RetornaTodosLotes() {
	lotes := l.lerDadosArquivo()
	lot := l.converteEmStruct(lotes)
	fmt.Println("---- Ler Dados: ",lot)

}

func  (l Lote) lerDadosArquivo() string {
    
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	nomeArquivo := "lote.txt"
	novoDir := filepath.Join(dir, "..", "..","domain","entities","lote")
	err = os.Chdir(novoDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return ""
	}

	// Lê todo o conteúdo do arquivo
	dados, err := ioutil.ReadFile(novoDir+ "/"+ nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return ""
	}

	return string(dados)
}

func (l Lote) converteEmStruct(dados string) []Lote {

	blocos := strings.Split(dados, "\n\n")
	loteMap := make(map[int]Lote)

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
					/*case "Produto":
						lote.Produto = campos[1]*/
					case "Quantidade":
						lote.Quantidade, _= strconv.Atoi(campos[1])
					case "DataValidade":
						lote.DataValidade  = campos[1]
					}
				}
			}

			idLote, _ := strconv.Atoi(lote.Id)
			if _, ok := loteMap[idLote]; !ok {
				loteMap[idLote] = lote
			}
		}
	}

	lotes := []Lote{}
	for _, lote := range loteMap {
		lotes = append(lotes, lote)
	}

	return lotes
}


