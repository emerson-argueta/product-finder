# Search for products by barcodes

## This is an API for searching products by GTIN-12 codes

### This project uses a custom [web-scraper](https://google.com) to create a google search for products. 

### &nbsp;
### &nbsp;

#### Before running the project, set the following environment variables or use a .env file
``` bash
PORT=8083
HOST=LOCALHOST
API_BASE_URL=/api
ENV=DEV
PRIVATE_KEY_PATH=./.keys/app.rsa
PUBLIC_KEY_PATH=./.keys/app.rsa.pub

DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mypass
DB_DB=productfinder
DB_SCHEMA=IDENTITY
DB_DIALECT=postgres
```

#### For API keys to work correctly, create a folder in the root directory of the project. Then Create private and public keys in that folder.
``` bash
mkdir .keys
cd .keys
openssl genrsa -out app.rsa 2048
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

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
You can register through this API with the POST endpoint at:
``` 
localhost:8083/api/register
```
For example:
``` HTTP
POST /api/register HTTP/1.1
Host: localhost:8083
Content-Type: application/json

{
    "email":"test@test.com",
    "password":"password"
}
```

This API has a GET endpoint at:
``` 
localhost:8083/api/productfinder/search
```
It requires an api_key value in the header and takes one query parameter as the GTIN-12 barcode:

``` HTTP
GET /api/productfinder/search?barcode=850008366079 HTTP/1.1
Host: localhost:8083
api_key: YOUR_API_KEY_OBTAINED_AFTER_REGISTERING
 
```



### &nbsp;
### &nbsp;

## Libraries used
* This project uses [goquery](https://github.com/PuerkitoBio/goquery) to convert the html search result into a jquery type object
* This project uses the [labstack echo](https://github.com/labstack/echo) router

