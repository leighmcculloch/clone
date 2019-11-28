build:
	go build

release:
	go run github.com/goreleaser/goreleaser --rm-dist
	./dist/clone_linux_amd64/clone -version
