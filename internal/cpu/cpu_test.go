package cpu

import (
	"fmt"
	"io/ioutil"

	"testing"
)

var files []string

func init() {
	const assets_path = "./assets/roms"
	dir, _ := ioutil.ReadDir(assets_path)
	for _, file := range dir {
		s := fmt.Sprintf("%s/%s", assets_path, file.Name())
		files = append(files, s)
	}

}

func TestCPUInit(t *testing.T) {
	for _, rom_path := range files {
		data, err := ioutil.ReadFile(rom_path)
		if err != nil {
			fmt.Printf("Fail to read ROM :%s\n", rom_path)
		}
		fmt.Printf("Loaded ROM :%s\nRom Size:%d Bytes\n", rom_path, uint16(len(data)))
		Init(data)
	}
}
