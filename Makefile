GOENV = GOOS=linux

bin:
	mkdir -p bin

bin/dbpopulator: bin cmd/dbpopulator/main.go
	$(GOENV) go build cmd/dbpopulator/main.go && mv main bin/dbpopulator

all: bin/dbpopulator