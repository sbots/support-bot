include linting.mk

local_run:
	export TEST_BOT_TOKEN="1", DOMAIN="1", PORT="8080",  SECRET_KEY="asd"
	go run main.go