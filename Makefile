compile:
	@echo "Removing previously built binaries"
	@rm -rf _output/bin || true
	@mkdir -p _output/bin
	@go build -o _output/bin/dw -v cmd/main.go

install:
	@echo "Creating symlink in ${GOPATH}/bin"
	@rm ${GOPATH}/bin/dw || true
	@ln -s `pwd`/_output/bin/dw ${GOPATH}/bin

test:
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./pkg/providers/
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./pkg/providers/libvirt
