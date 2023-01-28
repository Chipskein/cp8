package main

import (
	"chip8/internal/cpu"
	"io/ioutil"
	"log"
)

func main() {

	var rom_path = "./internal/__test_roms__/IBM_Logo.ch8"
	data, err := ioutil.ReadFile(rom_path)
	if err != nil {
		log.Printf("Fail to read ROM :%s\n", rom_path)
	}
	log.Printf("Loaded ROM :%s\nRom Size:%d Bytes\n", rom_path, uint16(len(data)))
	cpu.Init(data)
}
