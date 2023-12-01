build:
	@echo "Building binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o sendgrid-mail ./cmd/
	@echo "Done!"

build_local:
	@echo "Building binary..."
	env GOOS=darwin CGO_ENABLED=0 go build -o sendgrid-mail ./cmd/
	@echo "Done!"
