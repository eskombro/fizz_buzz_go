# fizz_buzz_go
Fizz-buzz go REST API

### GET /

`curl localhost:8000/`

### GET /health

`curl localhost:8000/health`

### POST /fizzbuzz

`curl -X POST localhost:8000/fizzbuzz -H 'Content-Type: application/json' -d '{"int1":3,"int2":5,"limit":30,"str1":"fizz","str2":"buzz"}'`

### GET /stats

`curl localhost:8000/stats`
