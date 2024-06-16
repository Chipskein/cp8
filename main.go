package main

import (
	c "cp8/cmd"
	"cp8/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCMD := cobra.Command{
		Use:   "cp8",
		Short: "cp8 is chip8 emulator",
		Long:  `cp8 is a simple chip8 emulator written in go.`,
		Run: func(cmd *cobra.Command, args []string) {
			c.Start(c.FilePath)
		},
	}
	rootCMD.Flags().BoolVarP(&config.DEBUG, "debug", "d", config.DEBUG, "set debug mode")
	rootCMD.Flags().StringVarP(&c.FilePath, "rom", "r", c.FilePath, "path to the rom file")
	rootCMD.MarkFlagRequired("rom")
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
