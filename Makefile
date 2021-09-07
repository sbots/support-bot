include linting.mk

app_name = "support-bot"
heroku_app_name = "support-bot-dev"

local_run:
	go run main.go

deps:
	go mod tidy
	go mod vendor

unit-test:
	go test -race -count=1 -v -cover ./...

build:
	docker build . --label $(app_name)

push:
	heroku container:push web -a $(heroku_app_name)

deploy:
	heroku container:release web -a $(heroku_app_name)

release: build push deploy