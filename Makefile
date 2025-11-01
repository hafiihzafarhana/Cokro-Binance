run:
	swag init -g cmd/app/main.go --output docs --parseDependency --parseInternal
	go run cmd/app/main.go

coveragetest:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
