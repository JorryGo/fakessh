Fake ssh for assholes.\
Our world has assholes which brute force ssh passwords on 22 port.\
This daemon sends "fake success" response for them.

```go get github.com/gliderlabs/ssh```\
```go build main.go```

You can find connection log in `connections.log` file