M = $(shell printf "\033[34;1m▶\033[0m")

build: schema ; $(info $(M) Building project...)
	go build -o manage

clean: ; $(info $(M) Removing generated files... )
	$(RM) api/schema/bindata.go

install: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	go mod vendor

schema: install clean ; $(info $(M) Embedding schema files into binary...)
	go generate ./...
	go run manage.go print-schema > api/schema.graphql

setup: ; $(info $(M) Fetching github.com/jteeuwen/go-bindata...)
	go get github.com/jteeuwen/go-bindata/...

server: schema ; $(info $(M) Starting development server...)
	go run manage.go runserver

image: ; $(info $(M) Building application image...)
	docker build -t inkster .

migrate: ; $(info $(M) Migrating database...)
	go run migrations/*.go up

test: schema migrate ; $(info $(M) Testing application...)
	GOCACHE=off go test ./... -p 1

test-update: schema migrate ; $(info $(M) Updating snapshots...)
	UPDATE_SNAPSHOTS=1 GOCACHE=off go test ./... -p 1

.PHONY: build clean install image schema setup server
