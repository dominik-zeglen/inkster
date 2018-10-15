M = $(shell printf "\033[34;1m▶\033[0m")

build: schema ; $(info $(M) Building project...)
	go build -o manage

clean: ; $(info $(M) [TODO] Removing generated files... )
	$(RM) api/schema/bindata.go

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	dep ensure

schema: dep clean ; $(info $(M) Embedding schema files into binary...)
	go-bindata \
		-ignore=\.snapshots \
		-ignore=\.go \
		-ignore=\.test \
		-pkg=schema \
		-o=api/schema/bindata.go \
		./api/...
	go run manage.go print-schema > api/schema.graphql

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep

server: schema ; $(info $(M) Starting development server...)
	go run manage.go runserver

image: ; $(info $(M) Building application image...)
	docker build -t inkster .

.PHONY: build clean dep image schema setup server
