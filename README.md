# HTTP API Server in Golang

[![codecov](https://codecov.io/gh/TonyPath/apiserver/branch/master/graph/badge.svg?token=H5CCXGQISL)](https://codecov.io/gh/TonyPath/apiserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/TonyPath/apiserver)](https://goreportcard.com/report/github.com/TonyPath/apiserver)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/TonyPath/apiserver)](https://github.com/TonyPath/apiserver)
[![GoDoc](https://godoc.org/github.com/TonyPath/apiserver?status.svg)](https://godoc.org/github.com/TonyPath/apiserver)
 
HTTP API Server in Go

## Examples

```go
func main() {
    r := NewRouter()
	r.Handle(http.MethodGet, "/test", func(w http.ResponseWriter, r *http.Request) {
        _, _ = w.Write([]byte(`test`))
    })
	
    apiSrv, _ := New(r, WithPort(8081))
    _ = apiSrv.Run(ctx)
}
```