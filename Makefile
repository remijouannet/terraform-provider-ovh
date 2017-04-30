build:
	go get ./...
	go build -o terraform-provider-ovh .
deps:
	go install github.com/hashicorp/terraform
install:
	go install .
test:
	go test -v ovh/*
