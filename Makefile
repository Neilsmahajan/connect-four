build: 
	@echo "Building the project..."
	go build -o ./bin/connect-four ./cmd/app/main.go

clean:
	@echo "Cleaning up..."
	rm -rf ./bin

run: build
	@echo "Running the project..."
	./bin/connect-four
