package cpu

import (
	"chip8/internal/cpu"
	"fmt"
	"io/ioutil"

	"testing"
)

func TestCpu(t *testing.T) {
	const assets_path = "./assets/roms"
	dir, _ := ioutil.ReadDir(assets_path)
	for _, file := range dir {
		s := fmt.Sprintf("%s/%s", assets_path, file.Name())
		fmt.Printf("Test ROM :%s\n", s)
		romData, _ := ioutil.ReadFile(s)
		cpu.Init(romData)
	}

}
