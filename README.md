## Go Deadlock

This project demonstrates a very simple deadlock example in Go.

## To Run

1. Start the HTTP server

```bash
go run app.go
```

2. Send concurrent requests using [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) or [bombardier](https://github.com/codesenberg/bombardier)

```bash
# ab https://httpd.apache.org/docs/2.4/programs/ab.html
ab -n 100 -c 50 'http://localhost:6060/'

# bombardier https://github.com/codesenberg/bombardier
bombardier -n 100 -c 50 -l http://localhost:6060
```
