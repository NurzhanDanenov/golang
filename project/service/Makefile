generate-swagger-user:
	go generate ./...
	docker run --rm -it  \
		-u $(shell id -u):$(shell id -g) \
		-e GOPATH=$(shell go env GOPATH):/go \
		-e GOCACHE=/tmp \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger:0.30.4 \
		generate spec -c ./cmd/user --scan-models -c ./internal/user -o ./swagger/OpenAPI/user.rest.swagger.json

generate-user-swagger:
	swagger generate spec --scan-models -c ./internal/user -c ./cmd/user -o ./swagger/OpenAPI/user.rest.swagger.json