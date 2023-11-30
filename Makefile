build:
	@echo "Building binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o sendgrid-mail ./cmd/
	@echo "Done!"
