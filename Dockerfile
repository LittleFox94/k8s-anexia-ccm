FROM golang:1.16 as builder
WORKDIR /go/src/github.com/github.com/anexia-it/anxcloud-cloud-controller-manager
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ccm

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/github.com/anexia-it/anxcloud-cloud-controller-manager/ccm .
CMD /app/ccm
