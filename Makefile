M = $(shell printf "\033[34;1m▶\033[0m")

build: schema ; $(info $(M) Building project...)
	go build -o manage

clean: ; $(info $(M) Removing generated files... )
	$(RM) api/schema/bindata.go

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	dep ensure

schema: dep clean ; $(info $(M) Embedding schema files into binary...)
	go generate ./...
	go run manage.go print-schema > api/schema.graphql

setup: ; $(info $(M) Fetching github.com/golang/dep and github.com/jteeuwen/go-bindata...)
	go get github.com/golang/dep/cmd/dep
	go get github.com/jteeuwen/go-bindata/...

server: schema ; $(info $(M) Starting development server...)
	go run manage.go runserver

image: ; $(info $(M) Building application image...)
	docker build -t inkster .

.PHONY: build clean dep image schema setup server
