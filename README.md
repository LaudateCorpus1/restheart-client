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
  ObjectType:    "object",
  RequestMethod: "GET",
}

err := rc.Call()
if err != nil {
  fmt.Println(err)
}
```

Client DELETE
```go
rc := restheart.Client{
  ObjectName:    "example",
  ObjectType:    "object",
  RequestMethod: "DELETE",
}

err := rc.Call()
if err != nil {
  fmt.Println(err)
}
```

Client PUT
```go

type Example struct{
  Name string
  IP string
}

example := Example{
  Name: "i-d987sdf98sdf09",
  IP: "10.0.0.1",
}

writeBuffer := new(bytes.Buffer)
json.NewEncoder(writeBuffer).Encode(example)

rc := restheart.Client{
  ObjectName:    "example",
  ObjectType:    "object",
  RequestMethod: "PUT",
  RequestPaload: writeBuffer.String()
}

err := rc.Call()
if err != nil {
  fmt.Println(err)
}
```

### TODO
* Add better error handling for Client
* Configuration to be read via environment variables
* Configuration to be settable via application that implements it
