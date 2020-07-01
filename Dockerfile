FROM golang:1.14-alpine AS build
WORKDIR /go/src/app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . ./
RUN go build -o /go/bin/imss /cmd/imss/imss.go

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/imss /
ENTRYPOINT ["/imss"]