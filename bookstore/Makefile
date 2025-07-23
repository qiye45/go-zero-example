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
