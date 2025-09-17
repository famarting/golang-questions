# Calculator API

Write a HTTP API that does arithmetic operations on two floats. The API should be able to perform the following operations at the given paths:

```
/add (+)
/subtract (-)
/multiply (*)
/divide (/)
```

Clients invoke the API using HTTP POST requests with the following JSON payload:

```
{
  "number_1": 1,
  "number_2": 2
}
```

Responses should be in the following format:

```
{
  "result": 3.0
}
```

e.g.

```
$ curl -X POST -H "Content-Type: application/json" -d '{"number_1": 1, "number_2": 2}' http://localhost:8080/add
{"result": 3.0}
```

Clients are also able to GET the previous result at /result. e.g.

```
$ curl http://localhost:8080/result
{"result": 3.0}
```

If the request body is not valid JSON or does not adhere to the above format, the API should return a 400 Bad Request response.
Any other Method (e.g. GET, PUT, DELETE) should return a 405 Method Not Allowed response.
If the requested operation is not supported, the API should return a 404.
Responses should always be valid JSON.