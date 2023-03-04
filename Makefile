deps:
	#go mod tidy \
    go get github.com/lucor/fyne-cross/cmd/fyne-cross \
#	go install github.com/fyne-io/fyne-cross@latest \
#	go get github.com/fyne-io/fyne-cross \

build:
	fyne-cross windows -arch=*

fmt:
	go fmt ./...

vet:
	go vet ./...

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)