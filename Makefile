all: db/up gen

gen/api: oapi.yml module/router/api/v1/events.yml
	@echo Generating OAPI
	@cd module/router/api/v1 && oapi-codegen -config ../../../../oapi.yml -package eventsv1 events.yml > events.gen.go

db/up: ./module/router/repo/migrations/*
	@echo Migrating DB
	@cd module/router/repo && migrate -database postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable -path ./migrations up

gen/db:
	@echo Generating PgGen
	pggen gen go --query-glob module/router/repo/queries.sql --postgres-connection postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable --go-type 'text=string' --go-type '_text=[]string' --go-type 'uuid=string' --go-type '_uuid=[]string' --go-type 'timestamptz=time.Time'

gen: gen/db gen/api
	@go mod tidy
	
run/mock:
	@echo Run Mock

run/demo:
	@echo Run Demo

	