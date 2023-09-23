package main

import (
	"chip8/internal/cpu"
	"log"
	"os"
)

func main() {
	var rom_path = "./internal/__test_roms__/IBM_Logo.ch8"
	data, err := os.ReadFile(rom_path)
	if err != nil {
		log.Fatalf("Fail to read ROM :%s\n", rom_path)
	}
	log.Printf("Loaded ROM :%s\nRom Size:%d Bytes\n", rom_path, uint16(len(data)))
	cpu.Init(data)
}
