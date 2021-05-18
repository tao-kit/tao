version := $(shell /bin/date "+%Y-%m-%d %H:%M")

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" taoctl.go
	$(if $(shell command -v upx), upx taoctl)
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o taoctl-darwin taoctl.go
	$(if $(shell command -v upx), upx taoctl-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o taoctl.exe taoctl.go
	$(if $(shell command -v upx), upx taoctl.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o taoctl-linux taoctl.go
	$(if $(shell command -v upx), upx taoctl-linux)
