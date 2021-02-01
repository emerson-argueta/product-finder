# Search for products by barcodes

## This is an API for searching products by GTIN-12 codes

### This project uses a custom [web-scraper](https://google.com) to create a google search for products. 

### &nbsp;
### &nbsp;

#### Before running the project, set the following environment variables
``` bash
HOST=LOCALHOST
PORT= 8083
HOST= LOCALHOST
API_BASE_URL=/api
ENV= DEV
SECRET=YOUR_SECRET_USED_FOR_JWT_TOKENS
```
#### Alternatively, you can use a .env file with the above values

### &nbsp;
### &nbsp;

## Run Project
```bash
# run tests
make test
# starts server
make run
```
### &nbsp;

## How it works
This API has one GET endpoint at:
``` 
localhost:8083/api/productfinder/search
```
It takes one query parameter as the GTIN-12 barcode:
```
localhost:8083/api/productfinder/search?barcode=850008366079
```


### &nbsp;
### &nbsp;

## Libraries used
* This project uses [goquery](https://github.com/PuerkitoBio/goquery) to convert the html search result into a jquery type object
* This project uses the [labstack echo](https://github.com/labstack/echo) router

