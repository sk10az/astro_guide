FROM golang:latest

COPY ./ ./
RUN make deps
RUN make build
CMD ["./main"]