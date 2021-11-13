.PHONY: build clean deploy gomodgen run-local

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/shifts shifts/deliveries/lambda/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/profile profile/deliveries/lambda/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
offline:
	sls offline start --useDocker --printOutput
local:
	CMDPORT=8005 TABLE_NAME=example-shifts go run cmd/server/main.go