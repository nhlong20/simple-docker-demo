# ltudweb-btcn-03-api

- port: 8080
- [GET] /v1/courses
- [POST] /v1/courses

## Run backend in simple mode 
- (no needed if you run docker-compose)
```bash
#!/usr/bin/env bash
echo "Downloading packages"
go mod download
echo "Compiling"

# build app in linux environment build image
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
```