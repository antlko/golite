install:
	npm install --prefix ./ui
	cp .env ui/.env

build-web: install
	npm run build --prefix ./ui

build: build-web
	go build .

run-web:
	npm run dev --prefix ./ui