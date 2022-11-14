package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var chip8Cmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := chip8Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
