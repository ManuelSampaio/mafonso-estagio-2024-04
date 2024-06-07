package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Cli() {
	comandoPrincipal := &cobra.Command{Use: "ACME"}

	expedirProduto := &cobra.Command{
		Use:   "expedirProduto [idProduto] [quantidade] ",
		Short: "expedir um determinado produto",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please specify a task to add")
				return
			}
			task := args[0]
			// Implementar a lógica para adicionar a tarefa aqui
			fmt.Printf("Adicionando tarefa: %s\n", task)
		},
	}

	comandoPrincipal.AddCommand(expedirProduto)
}
