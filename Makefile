.SILENT: new up up-to down down-to reset version status db-up db-down help jet tailwind templ air port kill

DSN = ${DRIVER}://${USER}:${PASSWORD}@tcp(${HOST}:${DB_PORT})/${DB}

include .env

install:
	@npm i -D tailwindcss
	@npm i -D daisyui@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/air-verse/air@latest
	@go install github.com/go-jet/jet/v2/cmd/jet@latest
	@go mod tidy

gen:
	@jet -source=mysql -dsn=${DSN} -path=./database

tailwind:
	@npx tailwindcss -i ./views/static/css/app.css -o ./views/static/css/index.css --minify --watch

templ:
	@templ generate --watch

tt:
	@make -j2 tailwind templ
	
air:
	@air

port:
	@netstat -ano | findstr :3000

kill:
	@taskkill /PID ${pid} /F