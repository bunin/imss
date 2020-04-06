PHONY: proto build

protoDeps := --proto_path=. --proto_path=${GOPATH}/src/github.com
h := CGO_ENABLED=0 GOOS=windows go build -o
t := -ldflags="-s -w" ./cmd/imss/imss.go

proto:
	protoc ${protoDeps} --go_out=,paths=source_relative:. --doc_out=. --doc_opt=markdown,types.md data/*.proto

build:
	${h} ./build/imss.exe ${t}
	GOARCH=386 ${h} ./build/imss_x86.exe ${t}
	upx ./build/*.exe