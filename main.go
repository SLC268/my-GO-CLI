package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "My GO CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("It Works!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
