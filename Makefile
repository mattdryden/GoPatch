all: run

run:
	@echo "Running..."
	go run *.go

build:
	@echo "Building binary..."
	go build

clean:
	@echo "Removing old binary..."
	rm -rf patches