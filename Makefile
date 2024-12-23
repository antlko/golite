# Install dependencies and copying envs
install:
	npm install --prefix ./ui
	cp .env ui/.env

build-web: install
	npm run build --prefix ./ui

# Build full app
build: build-web
	go build .

# Run with frontend
run: build
	./golite

# Build & Run backend only
run-be:
	go build .
	./golite dev

# Run UI only do dev purpose
run-ui:
	npm run dev --prefix ./ui