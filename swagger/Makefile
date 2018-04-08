BIN_DIR      :=      $(GOPATH)/bin
GOSWAGGER    :=      $(BIN_DIR)/swagger

.PHONY: all # Build the project
all: clean generate

.PHONY: fetch # Fetch new swagger.json from API Server
fetch:
	curl -o swagger.json "https://developers.zomato.com/swagger.json"

.PHONY: clean # Clean the previously generated files
clean:
	rm -rf zomato/ models/
	go clean -i

.PHONY: generate # Generate client library from swagger.json
generate: $(GOSWAGGER)
	swagger generate client -f=swagger.json -c=zomato

$(GOSWAGGER):
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger --install &> /dev/null

.PHONY: help # Show the help
help:
	@echo "Usage: make command \n\ncommands:"
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 : \2/' | expand -t20