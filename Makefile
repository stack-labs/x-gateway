NAME=x-gateway
IMAGE_NAME=docker.pkg.github.com/micro-in-cn/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=github.com/micro-in-cn/x-gateway/cmd
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG) -X $(GIT_IMPORT).BuildDate=$(BUILD_DATE)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)

all: build

vendor:
	go mod vendor

build:
	go build -a -installsuffix cgo -ldflags "-w ${LDFLAGS}" -o $(NAME) ./*.go

buildw:
	go build -a -installsuffix cgo -ldflags "-w ${LDFLAGS}" -o $(NAME).exe ./*.go

docker:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):latest

vet:
	go vet ./...

test: vet
	go test -v ./...

clean:
	rm -rf ./x-gateway

run_api:
	x-gateway --registry=$(registry) --transport=$(transport) api

run_web:
	x-gateway --registry=$(registry) --transport=$(transport) web

.PHONY: run_web run_api buildw build clean vet test docker
