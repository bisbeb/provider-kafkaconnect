FROM golang:1.23 AS builder

WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY apis/ apis/
COPY internal/ internal/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider cmd/provider/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/provider .
USER 65532:65532

ENTRYPOINT ["/provider"]
