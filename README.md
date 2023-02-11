# Aquafarm Management System

## Getting started

This is prototype of aquafarm management system

## Prerequisities
1. Docker
2. Golang >=1.19

## How To Run
1. Create env file
    - Create the env file by copy the example
        ```sh
        cp .env.example .env
        ```
    - Write down the value of env files. You can let the data as default except `JWT_PUBLIC_KEY`
2. Run the mysql by
    ```sh
    make db_init
    ```
3. Migrate the table by
    ```sh
    make db_migrate
    ```
4. Run the service by
    ```sh
    make service
    ```
If you want to test the code, run
```sh
make test
```

## API Docs

Note: Some of the endpoints need JWT. Please generate it using JWT Key (`JWT_PUBLIC_KEY`) and Algorithm (`JWT_ALG`) written in env file with payload format

```json
{
  "sub": "1234567890",
  "iat": 1516239022,
  "user": {
    "id": 1,
    "email": "john.doe@gmail.com"
  }
}
```
Example JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNTE2MjM5MDIyLCJ1c2VyIjp7ImlkIjoxLCJlbWFpbCI6ImpvaG4uZG9lQGdtYWlsLmNvbSJ9fQ.3HQk7dAYLBWe6LbauWUFz178rJcxa6agOHSL_lBUHdI`

### Get All Paginated Farms (GET /v1/farms)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Query Params
1. `page` = the current page
2. `limit` = number of data each page
3. `sorts` = list of sort data. You can sort it by
    - `-createdDate`: sort by created_at DESC
    - `createdDate`: sort by created_at ASC
    - `-code`: sort by code DESC
    - `code`: sort by code ASC
example: `?sorts=createdDate,-code`
4. `code` = filter data by code. Example: `?code=A1`
5. `village` = filter data by village. Example `?village=mustika`
6. `district` = filter data by district. Example: `?district=mustika`
7. `city` = filter data by city. Example: `?city=bekasi`
8. `province` = filter data by province. Example: `?province=Jawa`
9. `postalCode` = filter data by postalCode. Example: `?postalCode=17188`
10. `createdDateStart` = filter data whether greather than or equal
11. `createdDateEnd` = filter data whether less than or equal

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "count": 2,
        "currentPage": 1,
        "totalPage": 10,
        "limit": 100,
        "sorts": ["-createdDate", "code"],
        "items": [
            {
                "id": 1,
                "code": "A1",
                "description": "Description",
                ...
            }
        ]
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Create Farm (POST /v1/farms)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Body
{
    "code": "A1",
    "description": "description",
    "address": "Jln. Address",
    "village": "Pasir Gunung Selatan",
    "district": "Cimanggis",
    "city": "Kota Depok",
    "postalCode": "16451",
    "latitude": "-66.6666",
    "longitude": "170.899"
}

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "code": "A1",
        "description": "description",
        "address": "Jln. Address",
        "village": "Pasir Gunung Selatan",
        "district": "Cimanggis",
        "city": "Kota Depok",
        "postalCode": "16451",
        "latitude": "-66.6666",
        "longitude": "170.899",
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Delete Farm By ID (DELETE /v1/farms/:id)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

#### Response (200)
{
    "ok": true,
    "message": "Success",
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Get Farm By ID (GET /v1/farms/:id)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "code": "A1",
        "description": "description",
        "address": "Jln. Address",
        "village": "Pasir Gunung Selatan",
        "district": "Cimanggis",
        "city": "Kota Depok",
        "postalCode": "16451",
        "latitude": "-66.6666",
        "longitude": "170.899",
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Upsert Farm By Code (PUT /v1/farms/:code)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Body
{
    "description": "description",
    "address": "Jln. Address",
    "village": "Pasir Gunung Selatan",
    "district": "Cimanggis",
    "city": "Kota Depok",
    "postalCode": "16451",
    "latitude": "-66.6666",
    "longitude": "170.899"
}

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "code": "A1",
        "description": "description",
        "address": "Jln. Address",
        "village": "Pasir Gunung Selatan",
        "district": "Cimanggis",
        "city": "Kota Depok",
        "postalCode": "16451",
        "latitude": "-66.6666",
        "longitude": "170.899",
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}



### Get All Paginated Ponds (GET /v1/ponds)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Query Params
1. `page` = the current page
2. `limit` = number of data each page
3. `sorts` = list of sort data. You can sort it by
    - `-createdDate`: sort by created_at DESC
    - `createdDate`: sort by created_at ASC
    - `-code`: sort by code DESC
    - `code`: sort by code ASC
example: `?sorts=createdDate,-code`
4. `code` = filter data by code. Example: `?code=A1`
5. `volumeStart` = filter data by volume whether greather than or equal
6. `volumeEnd` = filter data by volume whether less than or equal
7. `areaStart` = filter data by area whether greather than or equal
8. `areaEnd` = filter data by area whether less than or equal
9. `createdDateStart` = filter data whether greather than or equal
10. `createdDateEnd` = filter data whether less than or equal

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "count": 2,
        "currentPage": 1,
        "totalPage": 10,
        "limit": 100,
        "sorts": ["-createdDate", "code"],
        "items": [
            {
                "id": 1,
                "farmId": 1,
                "code": "A1"
                "description": "Description",
                ...
            }
        ]
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Create Pond (POST /v1/ponds)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Body
{
    "farmId": 1,
    "code": "A1",
    "description": "Description",
    "wide": 5.67,
    "long": 10.89,
    "depth": 99.7
}

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "farmId": 1,
        "code": "A1",
        "description": "Description",
        "wide": 5.67,
        "long": 10.89,
        "depth": 99.7
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Delete Pond By ID (DELETE /v1/ponds/:id)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

#### Response (200)
{
    "ok": true,
    "message": "Success",
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Get Pond By ID (GET /v1/ponds/:id)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "farmId": 1,
        "code": "A1",
        "description": "Description",
        "wide": 5.67,
        "long": 10.89,
        "depth": 99.7
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Upsert Pond By Code (GET /v1/ponds/:code)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

##### Body
{
    "farmId": 1,
    "code": "A1",
    "description": "Description",
    "wide": 5.67,
    "long": 10.89,
    "depth": 99.7
}

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 1,
        "farmId": 1,
        "code": "A1",
        "description": "Description",
        "wide": 5.67,
        "long": 10.89,
        "depth": 99.7
        ...
    }
}

#### Response (400)
{
    "ok": false,
    "message": "Bad Request",
    "error": "Bad Request"
}

#### Response (404)
{
    "ok": false,
    "message": "Record Not Found",
    "error": "Record Not Found"
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Upsert All Endpoint Reports (GET /v1/endpoints/reports)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`

#### Response (200)
{
    "ok": true,
    "message": "Success",
    "data": {
       "GET /v1/ponds": {
            "count": 15,
            "uniqueUserAgent": 7,
            "uniqueIpAddress": 4,
            "requestTimeAverage": 6.77
       },
       "GET /v1/ponds/:id": {
            "count": 50,
            "uniqueUserAgent": 7,
            "uniqueIpAddress": 4,
            "requestTimeAverage": 4.4
       }
    }
}

#### Response (500)
{
    "ok": false,
    "message": "Internal Server Error",
    "error": "Internal Server Error"
}

### Get Swagger Documentation (GET /v1/docs/*)
Note: The endpoint won't return in reports
#### Response (200)