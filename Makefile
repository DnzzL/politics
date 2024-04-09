.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch --proxy=http://localhost:8080

.PHONY: init-db
init-db:
	touch politics.sqlite && go run internal/migrations/init_dbs.go

.PHONY: dev
dev:
	go build -o ./tmp ./cmd/main.go && air

.PHONY: build
build:
	make templ-generate  && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && make tailwind-build && jet -source=sqlite -dsn="politics.sqlite" -schema=dvds -path=./.gen