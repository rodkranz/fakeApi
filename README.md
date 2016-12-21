# Fake API

You don't need to wait for backend delivery your api any more, you can simulate the api response with this simple fakeApi, 
you can continue developing your frontend without dependencies.

It is a simple way to mock your api response.

## Source ##

* FakeApi Source
* Version: 1.0.0
* License: ISC


## How to Compile it

 You can see the full available binary list [here](https://gobuild.io/rodkranz/fakeApi)
 or compile those files from different platforms in your owner computer.

## Download

   Download for [Mac OSx](http://tmpcode.com/fake-api/fake-api_darwin_amd64.tar.gz)

   Download for [Linux 386](http://tmpcode.com/fake-api/fake-api_linux_386.tar.gz)

   Download for [Linux amd64](http://tmpcode.com/fake-api/fake-api_linux_amd64.tar.gz)

   Download for [Windows 386](http://tmpcode.com/fake-api/fake-api_windows_386.tar.gz)

   Download for [Windows 64](http://tmpcode.com/fake-api/fake-api_windows_amd64.tar.gz)


## Requirements

* [GO Language](https://golang.org/doc/install)

#### Compiling to *Linux*

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o fakeApi main.go


#### Compiling to *MacOSX*

	$ env GOOS=darwin GOARCH=386 go build -o fakeApi main.go


#### Compiling to *Windows*

	$ env GOOS=windows GOARCH=386 go build -o fakeApi.exe main.go


## Seed File ##
In a folder named `json`, you need to have the **seed** (json files) that will represent your api, the server will read all files inside folder and load them.
Use the file name to define the *URL* of api.

**e.g.**: If file name is `api_account_signup.json` the url will be `/api/account/signup`.

**The file seed needs to follow this format**: the seed file needs to follow this rules, *method*_*status_code*: *response* (Response can be any format)
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

You can add more then one response in seeds file for the same method and different methods too, just follow the rule in seed. 

**E.g**: Seed file name is `api_account_user.json` 
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
when you are using multiple response and no specify the status code in your header request, the response will be random between data that you putting in your seed file.
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

**Request POST**: use the header `X-Response-Code` to specify the response that you want to receive.
```
curl -X POST -H "X-Response-Code: 400" "/api/account/user"
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
