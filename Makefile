build:
	go build -o bin/chip8 ./main.go
test-cpu:
	go test ./internal/cpu/cpu_test.go -v