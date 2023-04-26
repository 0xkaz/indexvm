



# getNodeID


```
curl --location --request POST 'http://127.0.0.1:9650/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.getNodeID"
}'
```



# isBootstrapped
```
curl --location --request POST 'http://127.0.0.1:9650/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.isBootstrapped"
}'
```
# getNetworkID

```
curl --location --request POST 'http://127.0.0.1:9650/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.getNetworkID"
}'
```


-------


# hello ( new handler )
```

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/Mx33iLJsP8KRk918vayFJFAfkWjrtyLy4TycX7cbK6U3zuqpf/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.hello"
}'

```


# test ( new handler )
```

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/Mx33iLJsP8KRk918vayFJFAfkWjrtyLy4TycX7cbK6U3zuqpf/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.test", 
    "params": {
        "message": "adsfasdf"
    }
}'

```

# setData ( new handler ) => wrong impl.
```

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/Mx33iLJsP8KRk918vayFJFAfkWjrtyLy4TycX7cbK6U3zuqpf/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.setData", 
    "params": {
        "message": "ttttt"
    }
}'

```

# getData ( new handler ) => wrong impl.
```

curl --location --request POST 'http://127.0.0.1:9652/ext/bc/Mx33iLJsP8KRk918vayFJFAfkWjrtyLy4TycX7cbK6U3zuqpf/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.getData", 
    "params": {
    }
}'

```





