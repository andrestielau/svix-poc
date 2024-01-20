all: db/up gen

run: docker-compose up -d

gen/api: oapi.yml app/router/api/v1/event.yml app/webhook/api/v1/webhook.yml
	@echo Generating OAPI
	@cd app/router/api/v1 && oapi-codegen -config ../../../../oapi.yml -package eventsv1 event.yml > event.gen.go
	@cd app/webhook/api/v1 && oapi-codegen -config ../../../../oapi.yml -package webhooksv1 webhook.yml > webhook.gen.go

gen/grpc: buf.gen.yaml app/router/grpc/v1/event.proto app/webhook/grpc/v1/webhook.proto
	@buf generate 

db/up: ./app/router/repo/migrations/*
	@echo Migrating DB
	@cd app/router/repo && migrate -database postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable -path ./migrations up

gen/db:
	@echo Generating PgGen
	pggen gen go --query-glob app/router/repo/queries.sql --postgres-connection postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable --go-type 'text=string' --go-type '_text=[]string' --go-type 'uuid=string' --go-type '_uuid=[]string' --go-type 'timestamptz=time.Time'

gen: gen/db gen/grpc gen/api
	@go mod tidy

install:
	npm install @bufbuild/protobuf @bufbuild/protoc-gen-es @bufbuild/buf -g

run/mock:
	@echo Run Mock

run/demo:
	@echo Run Demo

	