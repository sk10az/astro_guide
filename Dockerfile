FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN #make deps
RUN go install github.com/fyne-io/fyne-cross@latest
RUN go get github.com/fyne-io/fyne-cross
RUN go mod tidy
RUN make build
CMD ["./astro_guide"]