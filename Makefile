TESTPATH=go-employees/pkg/services
IMAGE_NAME=go-empl
CONTAINER_NAME=go-employee-cont

.PHONY: proto
proto:
	protoc ./pkg/pb/*.proto --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \

.PHONY: test
test:
	go test -timeout 30s -v -run TestEmployeeTestSuite $(TESTPATH)

.PHONY: run
run:
	go run ./cmd/main.go

build:
	docker build . -t $(IMAGE_NAME)

up: build
	docker run --name $(CONTAINER_NAME) $(IMAGE_NAME) 