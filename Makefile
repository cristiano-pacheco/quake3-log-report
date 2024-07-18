# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install github.com/divan/expvarmon@latest
	go install github.com/rakyll/hey@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

# ==============================================================================
# Administration

run:
	go run ./cmd/cli/main.go

# ==============================================================================
# Running tests within the local computer
test-only:
	CGO_ENABLED=0 go test -count=1 ./...

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck -show verbose ./... 

test: test-only lint vuln-check