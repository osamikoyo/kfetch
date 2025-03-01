CC = go
OUTPUT = ./bin/kfetch

install:
	$(CC) install .
build:
	$(CC) build -o $(OUTPUT) main.go 
clean:
	rm -rf $(OUTPUT)