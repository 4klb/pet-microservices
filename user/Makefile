.PHONY: gen-user
gen-user:
	protoc -I proto/user --go_out=proto user.proto

.PHONY: gen-userService
gen-userService:
	protoc -I proto --go_out=proto --go-grpc_out=proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto
	