# -- docker builds
IMAGE_SERVER ?= payment-payments-api-server:latest
build-server:
	docker build . --build-arg cmd=server -t $(IMAGE_SERVER)

# -- docker compose
start: build-server
	docker-compose up -d

stop:
	docker-compose stop

clean: stop
	docker system prune --all --volumes

# -- dev tools
lint:
	golangci-lint run

# -- backend seeders
install-seeders-tools:
	brew install curl
	brew install jq

ADMIN_TOKEN=$(shell curl -s -X POST http://localhost:5000/users/login -H 'content-type: application/json' -d '{"email": "admin@example.com","password": "admin123"}' | jq -r .data.token)

upload-fees:
	@curl -X POST \
      http://localhost:5000/fees/massive-upload \
      -H 'authorization: $(ADMIN_TOKEN)' \
      -H 'content-type: multipart/form-data' \
      -F file=@./internal/api/testing/assets/fees.xlsx
