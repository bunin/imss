set TAIL=-ldflags="-s -w" .\cmd\imss\imss.go

set CGO_ENABLED=0
go build -o imss.exe %TAIL%

set GOARCH=386
go build -o imss-i386.exe %TAIL%
