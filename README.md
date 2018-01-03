# restheart-go
This package is a simple RESTHeart client.

Client struct provides a method which constructs an HTTP/S request and makes a call to a RESTHeart endpoint.  

Configuration struct provides a method to configure the Client with credentials and HTTP/S RESTHeart endpoint.

### Usage

Import
```go
import restheart "github.vianttech.com/techops/restheart-go"
```

Client GET
```go
rc := restheart.Client{
  ObjectName:    "example",
  ObjectType:    ,
  RequestMethod: "GET",
}

err := rc.Call()
if err != nil {
  return "", err
}
```
