ifeq ($(OS), Windows_NT)
	TARGET = main.exe
else
	TARGET = main
endif

run: build
	@$(TARGET)
build:
	@go build -o $(TARGET) ./cmd/
