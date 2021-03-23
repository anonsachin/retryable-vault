run:
	go run main.go

vault:
	vault server -dev -dev-root-token-id=myroot

call:
	curl http://localhost:8080/vault

kv:
	curl -X POST -d @kv.json http://localhost:8080/vault

unittests:
	go test ./pkg/send/*.go -tags=unittest
	go test ./pkg/handler/*.go -tags=unittest