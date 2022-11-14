build:
	go build -o bin/chip8 ./main.go
test-cpu:
	go test ./internal/cpu/cpu_test.go -v
test-screen:
	go test ./internal/screen/screen_test.go -v
