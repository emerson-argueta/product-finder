
GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

all: test build

test: 
	godotenv -f .env $(GOTEST) ./modules/productfinder/infrastructure/webscraper -cover -v
	godotenv -f .env $(GOTEST) ./modules/identity/infrastructure/http/routes -cover -v
run:
	godotenv -f .env $(GORUN) ./cmd/*

test-database:
	cd ./.test && docker-compose down -v
	cd ./.test && docker-compose up