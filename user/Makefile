GO := @go
NODEMON := @nodemon

# import deploy config
# You can change the default deploy config with `make cnf="deploy_special.env" release`
env ?= .env
include $(env)
export $(shell sed 's/=.*//' $(env))

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub

goinstall:
	${GO} get .

godev:
	${NODEMON} --watch './**/*.go' --signal SIGTERM --exec 'go' run ./main.go

goprod:
	${GO} build -o main .

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...

$(eval VERSION:=$(shell git rev-parse --short HEAD))
VERSION=i7aa9j5

# DOCKER TASKS
# Build the container
build: ## Build the container
	docker build -t $(APP_NAME) .

build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

release: build-nc publish ## Make a release by building and publishing the `{version}` ans `latest` tagged containers to Repository

# Docker publish
publish: publish-latest publish-version ## Publish the `{version}` ans `latest` tagged containers to Repository

publish-latest: tag-latest ## Publish the `latest` taged container to Repository
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):latest

publish-version: tag-version ## Publish the `{version}` taged container to Repository
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

# Docker tagging
tag: tag-latest tag-version ## Generate container tags for the `{version}` ans `latest` tags

tag-latest: ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):latest

tag-version: ## Generate container `latest` tag
	@echo 'create tag $(VERSION)'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):$(VERSION)