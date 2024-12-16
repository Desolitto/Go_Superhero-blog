
all: server_run

logo: 
	go build cmd/logocreater/logo.go

logo_run:
	go run cmd/logocreater/logo.go

server: logo_run
	@go build cmd/server/server.go

.PHONY: server_run

server_run: logo_run
	@go run cmd/server/server.go || true

clean:
	rm -rf logo, server

