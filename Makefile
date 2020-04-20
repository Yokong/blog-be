.PHONY: build docker clean help ck

all: build docker clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=prod -o blog-be
	@echo "------ build go success ------"

docker:
	docker build -t blog-be .
	docker tag blog-be yokooll/blog-be
	docker push yokooll/blog-be
	@echo "------ docker push success ------"

clean:
	@rm -rf blog-be
	@go clean -i .
	@echo "------ clean done ------"
	@docker images

ck:
	@docker images|grep none|awk '{print $3}'|xargs docker rmi -f

help:
	@echo "make: compile packages and dependencies"
	@echo "make docker: build,tag,push docker image"
	@echo "make clean: remove object files and cached files"
	@echo "make doc: swag init"
	@echo "make ck: clean docker images"
	@echo "make genpb: generate pb files"