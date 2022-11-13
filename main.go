package main

import (
	"chip8/internal/cpu"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("./test.ch8")
	if err != nil {
		log.Panic("Could not load test")
	}
	cpu.Init(data)

}
