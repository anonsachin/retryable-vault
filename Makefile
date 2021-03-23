run:
	go run main.go

vault:
	docker-compose -f dependency/docker-compose.yml up -d

vault-down:
	docker-compose -f dependency/docker-compose.yml down -v

call:
	curl http://localhost:8080/vault

kv:
	curl -X POST -d @kv.json http://localhost:8080/vault

tests:
	go test ./pkg/send/*.go
	go test ./pkg/handler/*.go

tests-with-dependency:
	go test ./pkg/tests/*.go

all-tests: tests vault tests-with-dependency vault-down