# Fake API


I don't need to wait for backend to develop your frontend page, you can simulate the api
simple way to show json with same the response of api.

## Source ##

* FakeApi Source
* Version: 1.0.0
* License: ISC


## Seed File ##
Create a folder's name `json`, The all files inside this folder will be loaded in seed to create urls, defined the name of file to represent the `URL`.

**e.g.**: If file name is `api_account_signup.json` the url will be `/api/account/signup`.

**The seed needs to follow this format**: format needs to follow this rules, *method*_*status_code*: *response* (the response can be format)
```
{
    "[METHOD]_[STATUS_CODE]": [RESPONSE]
}
```

**Example seed format**: 
```
{
	"POST_200": {
        "response": "Post Request with status code 200",
        "statusCode: 200,
	}
}        
```

**Example request**:
```
curl -X POST "http://localhost:9090/api/account/signup"
```

**Response will be**: 
```
{
	"response": "POST Request with status code 200",
	"statusCode: 200,
}
```


## Multiples Response for seed ##

You can add more then one response in seeds file, just follow the rule in seed. 

Seed file name is `api_account_user.json` 

**E.g**:
```
{
   "GET_200": {
        "response": "GET Request with status code 200",
        "statusCode: 200,
	},
    "POST_400": {
        "response": "POST Request with status code 400",
        "statusCode: 400,
        "error": {
        	"email": "email has invalid format",
            "time": "the date is invalid"
        }
	}
    "POST_200": {
        "response": "Post Request with status code 200",
        "statusCode: 200,
	}
}
```

**Request POST**: 
when you use multiple status code to response without header to specific status code to response the response will be always random between status codes available for method resquest.

```
curl -X POST "/api/account/user"
```

**Response POST Dynamic**: 
```
{
	"response": "POST Request with status code 400",
	"statusCode: 400,
	"error": {
		"email": "email has invalid format",
		"time": "the date is invalid"
	}
}
```
or 
```
{
	"response": "Post Request with status code 200",
	"statusCode: 200,
}
```

**Request POST**: use the header with name `X-Requested-Code` to specify the response that you want to receive.
```
curl -X POST -H "X-Requested-Code: 400" "/api/account/user"
```


**Response POST specific**: 
```
{
	"response": "POST Request with status code 400",
	"statusCode: 400,
	"error": {
		"email": "email has invalid format",
		"time": "the date is invalid"
	}
}
```

OBS: By default the cross domain is always enabled.