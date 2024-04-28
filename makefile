ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_COMPOSE_FILE := docker-compose.yml
SWAGGER_COMPOSE_FILE := docker-compose_swagger.yml
APP_CONTAINER_NAME := template-api
SWAGGER_CONTAINER_NAME := swagger
TERRAFORM_CMD = docker run --rm -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -v  $(PWD)/iac/terraform:/terraform -w /terraform -it hashicorp/terraform:1.7.5 -chdir=/terraform/$(TF_ENV)/main
TERRAFORM_PRE_CMD = docker run --rm -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -v  $(PWD)/iac/terraform:/terraform -w /terraform -it hashicorp/terraform:1.7.5 -chdir=/terraform/$(TF_ENV)/pre_main

.PHONY: up
up: 
	docker-compose -f ${APP_COMPOSE_FILE} up -d
	cp ./doc/swagger/api.yml ./app/openapi/.

.PHONY: up-build
up-build: 
	docker-compose -f ${APP_COMPOSE_FILE} up -d --build

.PHONY: up-build-swagger
up-build-swagger: 
	docker-compose -f ${SWAGGER_COMPOSE_FILE} up -d --build
	cp ./doc/swagger/openapi.yml ./app/openapi/.

.PHONY: gen-openapi
gen-openapi: 
	rm -rf ./doc/swagger/openapi.yml
	rm -rf ./app/openapi/openapi.yml
	docker-compose -f ${SWAGGER_COMPOSE_FILE} up swagger-cli -d --build
	@sleep 2 # コピーするまでに作成が間に合わないので、スリープ
	cp ./doc/swagger/openapi.yml ./app/openapi/.

.PHONY: ui
ui: 
	docker-compose -f ${SWAGGER_COMPOSE_FILE} up swagger-ui -d --build

.PHONY: test
test: 
	docker container exec -it ${APP_CONTAINER_NAME} go test ./... -coverprofile=coverage.out
	docker container exec -it ${APP_CONTAINER_NAME} go tool cover -html=coverage.out -o coverage.html

.PHONY: wire
wire: 
	docker container exec -it ${APP_CONTAINER_NAME} go generate -x -tags wireinject ./wire/wire.go

.PHONY: oapi-codegen
oapi-codegen:
	docker container exec -it ${APP_CONTAINER_NAME} oapi-codegen -package openapi -generate types,server -o ./openapi/openapi.gen.go /src/app/openapi/openapi.yml

.PHONY: lint
lint: 
	docker container exec -it ${APP_CONTAINER_NAME} golangci-lint run

.PHONY: mock-gen
mock-gen: 
	docker container exec -it ${APP_CONTAINER_NAME} sh -c 'cd mock && go generate ./...'

.PHONY: init plan apply destroy

init:
	$(TERRAFORM_CMD) init

plan:
	$(TERRAFORM_CMD) plan

apply:
	$(TERRAFORM_CMD) apply

destroy:
	$(TERRAFORM_CMD) destroy

.PHONY: init-pre plan-pre apply-pre destroy-pre

init-pre:
	$(TERRAFORM_PRE_CMD) init

plan-pre:
	$(TERRAFORM_PRE_CMD) plan

apply-pre:
	$(TERRAFORM_PRE_CMD) apply

destroy-pre:
	$(TERRAFORM_PRE_CMD) destroy
