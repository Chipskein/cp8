package cmd

import (
	"chip8/internal/cpu"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var chip8Cmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var numArgs = len(args)
		if numArgs == 0 {
			log.Panicf("No ROM path passed as argument\n")
			os.Exit(1)
		}
		var rompath = args[0]
		log.Println("Iniciando Chip8 Machine", rompath)
		data, err := ioutil.ReadFile(rompath)
		if err != nil {
			log.Panicf("Could not Read File:%s\n Error: %s\n", rompath, err)
			os.Exit(1)
		}
		cpu.Init(data)
	},
}

func Execute() {
	if err := chip8Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
