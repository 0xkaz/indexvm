
# NOTE


```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc

go mod tidy
go mod vendor


go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```
protoc --go_out=. *.proto
protoc --go-grpc_out=. *.proto
```

etc



# reflection  

## how to install grpcurl 

```
brew install grpcurl
```

## reflection test 
```
grpcurl -plaintext localhost:3003 list 
```

# watch 

## install watch
```
brew install watch

```




# ref

https://grpc.io/docs/languages/go/quickstart/

# how to start dev under this directory

# install protoc

## macosx

```
brew install protobuf-c
```

## linux ( apt )

```
apt install -y protobuf-compiler
```

## check version

```
protoc --version

```


