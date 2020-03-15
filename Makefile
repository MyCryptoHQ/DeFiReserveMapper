.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/app app/main.go
clean:
	rm -rf ./bin ./vendor Gopkg.lock
deploy: clean build
	zip -r app.zip 'assets.json' 'bin/app'
