run:
	swag init -g cmd/app/main.go --output docs --parseDependency --parseInternal
	go run cmd/app/main.go
