# Initial go environment
.PHONY: env
env:
	go mod tidy

# Start api
.PHONY: api
api:
	go run api/bookstore.go -f api/etc/bookstore-api.yaml

# Start add rpc
.PHONY: add
add:
	go run rpc/add/add.go -f rpc/add/etc/add.yaml

.PHONY: check
check:
	go run rpc/check/check.go -f rpc/check/etc/check.yaml

.PHONY: docker-build
# 构建docker镜像

docker-build:
	docker build -t bookstore:latest .

.PHONY: docker-run-api
# 运行api服务容器

docker-run-api:
	docker run --rm -p 8888:8888 -v $(PWD)/api/etc:/app/api/etc bookstore:latest ./api_server -f api/etc/bookstore-api.yaml

.PHONY: docker-run-add
# 运行add rpc服务容器

docker-run-add:
	docker run --rm -p 8081:8081 -v $(PWD)/rpc/add/etc:/app/rpc/add/etc bookstore:latest ./add_rpc -f rpc/add/etc/add.yaml

.PHONY: docker-run-check
# 运行check rpc服务容器

docker-run-check:
	docker run --rm -p 8082:8082 -v $(PWD)/rpc/check/etc:/app/rpc/check/etc bookstore:latest ./check_rpc -f rpc/check/etc/check.yaml

.PHONY: compose-up
compose-up:
	docker compose -f deploy/docker-compose.yaml up -d

.PHONY: compose-down
compose-down:
	docker compose -f deploy/docker-compose.yaml down

.PHONY: helm-install
helm-install:
	cd deploy/helm && helm install bookstore .

.PHONY: helm-uninstall
helm-uninstall:
	cd deploy/helm && helm uninstall bookstore
