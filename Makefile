GO = go
GOENV = GOOS=linux

pkgs  = $(shell GOFLAGS=-mod=vendor $(GO) list ./... | grep -vE -e /vendor/ -e /pkg/swagger/)

bin:
	mkdir -p bin

clean:
	rm bin/*

bin/dbpopulator: bin cmd/dbpopulator/main.go pkg/*
	$(GOENV) $(GO) build -o bin/dbpopulator cmd/dbpopulator/main.go

bin/restserver: bin cmd/restserver/main.go pkg/*
	$(GOENV) $(GO) build -o bin/restserver cmd/restserver/main.go

#-------------------------
# Code generation
#-------------------------
.PHONY: generate

## Generate go code
generate:
	@echo "==> generating go code"
	GOFLAGS=-mod=vendor $(GO) generate $(pkgs)

#-------------------------
# Target: swagger.validate
#-------------------------
.PHONY: swagger.validate

swagger.validate:
	swagger validate pkg/swagger/swagger.yaml

#-------------------------
# Target: swagger.doc
#-------------------------
.PHONY: swagger.doc

swagger.doc:
	mkdir -p doc && docker run --rm -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yaml > doc/index.html

swagger-all: swagger.validate generate swagger.doc

all: swagger-all bin/dbpopulator bin/restserver

localrest: bin/restserver
	./bin/restserver 