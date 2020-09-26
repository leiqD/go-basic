BINARY=bin/demo
VERSION=$(shell git describe --abbrev=0 --tags 2> /dev/null || echo "0.1.0")
BUILD=$(shell git rev-parse HEAD 2> /dev/null || echo "undefined")
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

test: 
	go test -v -cover -covermode=atomic ./...

demo:
	go build -o ${BINARY} $(LDFLAGS) app/socket5/*.go


unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

mock:
	$(GOPATH)/bin/mockgen -source=./uc/INTERACTOR.go -destination=./implem/uc.mock/interactor.go -package=mock && \
    $(GOPATH)/bin/mockgen -source=./uc/HANDLER.go -destination=./implem/uc.mock/handler.go -package=mock

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint