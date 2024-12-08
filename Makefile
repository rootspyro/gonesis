APP=gonesis
GCO_ENABLED=0
GOOS=linux 
GOARCH=amd64


build:
	GCO_ENABLED=$(GCO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP) cmd/$(APP)/main.go
