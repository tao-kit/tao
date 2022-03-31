build:
	go build -ldflags="-s -w" taoctl.go
	$(if $(shell command -v upx), upx taoctl)
mac:
	GOOS=darwin go build -ldflags="-s -w" -o taoctl-darwin taoctl.go
	$(if $(shell command -v upx), upx taoctl-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -o taoctl.exe taoctl.go
	$(if $(shell command -v upx), upx taoctl.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -o taoctl-linux taoctl.go
	$(if $(shell command -v upx), upx taoctl-linux)
