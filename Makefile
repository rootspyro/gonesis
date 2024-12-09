APP=gonesis
GCO_ENABLED=0
GOOS=linux 
GOARCH=amd64
VERSION=1.0.0


build:
	GCO_ENABLED=$(GCO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.version=$(VERSION)" -o bin/$(APP) cmd/$(APP)/main.go
