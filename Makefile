.PHONY: fmt build clean deploy

fmt:
	cd app/ && go fmt
	cd pkg/ && go fmt ./...
	cd terraform/ && terraform fmt -diff
build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/app app/main.go
clean:
	rm -rf ./bin ./vendor Gopkg.lock
deploy: fmt clean build
	zip -r app.zip 'assets.json' 'bin/app'
