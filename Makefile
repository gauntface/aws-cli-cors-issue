clean:
	rm -rf .aws-sam \
	go mod tidy

build: clean
	sam build

serverdefault: build
	sam local start-api \
		--port 3000 \
		--docker-network focus-api_local-api-network

serverparams: build
	sam local start-api \
		--port 3000 \
		--parameter-overrides="GatewayCorsOrigins='http://localhost:3001,http://localhost:3002'" \
		--docker-network focus-api_local-api-network

curl:
	curl 'http://localhost:3000/hello' -X 'OPTIONS' -i