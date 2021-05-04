.PHONY: server client run clean
server:
	go build -o bin/meshd cmd/meshd/meshd.go

client:
	go build -o bin/meshdc cmd/meshd-client/client.go
 
run: server
	bin/meshd
 
clean:
	rm -rf bin/