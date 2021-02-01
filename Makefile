
GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

all: test build

test: 
	godotenv -f .env $(GOTEST) ./modules/productfinder/infrastructure/webscraper -cover -v
run:
	godotenv -f .env $(GORUN) ./cmd/*