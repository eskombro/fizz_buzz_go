# fizz_buzz_go

#### Fizz-buzz go REST API

The goal is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.



The server needs to be:

- Ready for production
- Easy to maintain by other developers



Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:

- Accept no parameter

- Return the parameters corresponding to the most used request, as well as the number of hits for this request

## Install

1. Install `docker` and `docker-compose`  
2. Run `git clone https://github.com/eskombro/fizz_buzz_go.git && cd fizz_buzz_go`
3. Run `docker-compose up -d`

## Implementation

Deployment built using **Docker**. Docker containers are:
- **Nginx** exposed at port 4242
- **Go REST API** (based on gorilla-mux)
- **MongoDB** for data persistence

## Test it

**CI** implemented with **github actions**. Every new 'push' will automatically run tests:

[See FizzBuzz github actions](https://github.com/eskombro/fizz_buzz_go/actions)

To manually run tests, use:  

`cd fizz_buzz_go/src && go test -v ./...`

API is available for test at:

http://51.75.23.195:4242/

## API Documentation

##### GET /

Home page, returns information about main endpoints

Example:  
`curl localhost:4242/`

Response:  
```"FIZZBUZZ: Use POST /fizzbuzz to execute, or GET /stats to see stats"```  

##### GET /health  

Health endpoint. Request to check if service is running  

Example:  
`curl localhost:4242/health`

Response:  
```"Running"```  

##### POST /fizzbuzz

Example:  
`curl -X POST localhost:4242/fizzbuzz -H 'Content-Type: application/json' -d '{"int1":3,"int2":5,"limit":30,"str1":"fizz","str2":"buzz"}'`  

Response:  
```"1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz, 16, 17, fizz, 19, buzz, fizz, 22, 23, fizz, buzz, 26, fizz, 28, 29, fizzbuzz"```  

##### GET /stats

Example:  
`curl localhost:4242/stats`

Response:  
```{"params":{"int1":3,"int2":5,"limit":30,"str1":"fizz","str2":"buzz"},"count":1}```
