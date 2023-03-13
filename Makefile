build-cmd:
	go build .
	sudo mv easypdf /usr/local/bin

cmd-first-build:
	go mod download
	go build .
