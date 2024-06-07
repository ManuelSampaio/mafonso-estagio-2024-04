package cli

import (
	"fmt"
	"os"
	"ACMELDA/interno/interface/servicos"
	"github.com/spf13/cobra"
)

func Cli() {
	var idProduto string
	var quantidade int
	comandoPrincipal := &cobra.Command{Use: "ACME"}
	expedirProduto := &cobra.Command{
		Use:   "expedirMercadoria [idProduto] [quantidade] ",
		Short: "expedir uma determinada mercadoria",
		Run: func(cmd *cobra.Command, args []string) {
			if idProduto == "" {
				fmt.Println("Por favor especifique o identificador do produto a expedir")
				return
			}

			if quantidade == 0 {
				fmt.Println("Por favor especifique a quantidade a expedir")
				return
			}

			fmt.Println("o id do produto: ", idProduto)
			fmt.Println("a quantidade: ", quantidade)
			servicos.ExpedirMercadoria(idProduto, quantidade)
		},
	}

	expedirProduto.Flags().StringVarP(&idProduto, "idProduto", "i", "", "identificador do produto a expedir")
	expedirProduto.Flags().IntVarP(&quantidade, "quantidade", "q", 0, "a quantidade do produto a expedir")

	comandoPrincipal.AddCommand(expedirProduto)

	if err := comandoPrincipal.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
