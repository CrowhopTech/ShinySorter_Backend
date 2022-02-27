GOENV = GOOS=linux

bin:
	mkdir -p bin

bin/dbpopulator: bin cmd/dbpopulator/main.go
	$(GOENV) go build cmd/dbpopulator/main.go && mv main bin/dbpopulator

bin/restserver: bin cmd/restserver/main.go
	$(GOENV) go build cmd/restserver/main.go && mv main bin/restserver

all: bin/dbpopulator bin/restserver