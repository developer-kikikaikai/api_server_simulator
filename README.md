# api_server_simulator
The simulator of api backend server

# What is this?

This is the api server for REST API frontend developer.  
You can POST endpoint and response which you want.

# How to use?
1. run server.
2. POST /endpoints with following JSON code.

```
{
  "endpoint": "string",
  "method": "string",
  "response": {
    "body": "string",
    "header": [
      {
        "type": "string",
        "value": "string"
      }
    ],
    "status": 200
  }
}
```

- You can choose status, endpoint and method.  
3. After PUT JSON code, you can use above endpoint. It will reply 

Please see swagger.yaml to know more API usage.

# How to run server?

1. Install golang and mongo DB.  

[MongoDB Download Center](https://www.mongodb.com/download-center?utm_source=manual&utm_campaign=download-mongodb-navbar-cta&utm_medium=docs)
[Golang install](https://golang.org/doc/install)

2. Install dep
[golang/dep](https://github.com/golang/dep)

3. Get this code.
`go get -v github.com/developer-kikikaikai/api_server_simulator`

4. Ensure related package.  

```
cd $GOPATH/src/github.com/developer-kikikaikai/api_server_simulator
dep ensure
```

5. Run main.go
`go run main.go` or `make`

TODO: Save configuration to DB.

# How to do test it?

Please use `go test -tags=testdb -v ./...`.  
* Please run with "-tags=testdb" option to use DB for test.

And if you run it on Linux, you can do the API test.  
Please call `./unittest.sh` in test directory.  

TODO: Almost REST API tests have already finished. But I only haven't created new endpoint test which defined specific HTTP headers. 
(The case if requested json data has headers.)

# How to run docker container

Please call `make container` on root directory.
* You can import api data to put the "import.json" file in docker dir
  The import.json is same format of GET endpoints request, as

```
[
{
  "endpoint": "string",
  "method": "string",
  "response": {
    "body": "string",
    "header": [
      {
        "type": "string",
        "value": "string"
      }
    ],
    "status": 200
  }
}
]
```

* Please install docker and ruby

```
sudo apt install -y docker ruby
```

# Others

License : Apache2.0 License. Please see LICENSE file.  
