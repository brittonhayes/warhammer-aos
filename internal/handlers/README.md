<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# handlers

```go
import "github.com/brittonhayes/warhammer-aos/internal/handlers"
```

package handlers contains all http request handlers for the API

## Index

- [type Handler](<#type-handler>)
- [type Response](<#type-response>)


## type Handler

Handler contains all routes for API request handling

```go
type Handler interface {
    AddRoute(path string, handlers ...fiber.Handler) *handle
}
```

## type Response

Response is the default API response format

```go
type Response struct {
    Count int         `json:"count"`
    Data  interface{} `json:"data"`
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
