



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

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2QqEkALmvXzNhYXXuTGV9STzeMZ98CfZWXqtFxetJ4AeCiSThq/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.hello"
}'

```





