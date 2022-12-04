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

image:
	docker build --rm --platform linux/amd64 -t kevinwan/taoctl:$(version) .
	docker tag kevinwan/taoctl:$(version) kevinwan/taoctl:latest
	docker push kevinwan/taoctl:$(version)
	docker push kevinwan/taoctl:latest
	docker build --rm --platform linux/arm64 -t kevinwan/taoctl:$(version)-arm64 .
	docker tag kevinwan/taoctl:$(version)-arm64 kevinwan/taoctl:latest-arm64
	docker push kevinwan/taoctl:$(version)-arm64
	docker push kevinwan/taoctl:latest-arm64
