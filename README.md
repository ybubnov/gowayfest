# Bench

```sh
docker run --rm --name ofp --network=host -it -v $HOME/go/src:/go/src -w /go/src/github.com/netrack/bench -m 128M --memory-swap=128M golang:1.10.4 go run server.go
```
