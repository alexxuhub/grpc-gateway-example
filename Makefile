.PHONY: help build proto  make file

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z0-9_/-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


server: ## start the grpc gateway server
	@go run main/main.go server

proto: ## gen grpc for proto
	@rm -rf ./pkg/xd
	@protoc -I ./api \
    		 --go_out=. \
             --go-grpc_out=. \
    		 api/*/*.proto
gateway: ## gen gateway
	@protoc -I ./api \
	         --grpc-gateway_out ./pkg/alex \
	         --grpc-gateway_opt logtostderr=true \
	         --grpc-gateway_opt paths=source_relative \
	         --grpc-gateway_opt generate_unbound_methods=true \
	         api/*/*.proto

openapi: ## gen swagger
	@protoc -I ./api \
              --openapiv2_out ./doc \
              --openapiv2_opt logtostderr=true \
               api/*/*.proto