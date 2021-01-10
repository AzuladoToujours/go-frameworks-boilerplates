### GO SIMPLE API

```shell
## Create main
touch main.go

## Create dependencies
go mod init [repository]

## Install gin framework
go get -u github.com/gin-gonic/gin

##Ways to test
go test ./platform/newsfeed/

go test ./...

go test -cover ./...

## Ways to run

make dev

docker build -t simple-api:1 .

docker run -it --rm -d --name simple-api -p 8080:8080/tcp simple-api:1
```