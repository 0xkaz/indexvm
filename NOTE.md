



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

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.hello"
}'

```


# test ( new handler )
```

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
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

curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
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

curl --location --request POST 'http://127.0.0.1:9652/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.getData", 
    "params": {
    }
}'

```



# Creates a new key in the default location

```
./build/index-cli create 
```


# network 

```
./build/index-cli network --endpoint http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3
```


# balance 

```
curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.balance", 
    "params": {
        "address": "index1rvzhmceq997zntgvravfagsks6w0ryud3rylh4cdvayry0dl97nsqrawg5"
    }
}'

```

```
curl --location --request POST 'http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3/rpc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"indexvm.balance", 
    "params": {
        "address": "index1l97kg5xvpxm0qvhyy6vuej2llg7ydlcp86euda6ewk3uuq4jks0q2rm5qx"
    }
}'

```

# transfer 

```
./build/index-cli transfer index1rvzhmceq997zntgvravfagsks6w0ryud3rylh4cdvayry0dl97nsqrawg5 15 --endpoint http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3 --private-key-file .index-cli.pk
```

# network 

```
./build/index-cli watch --endpoint http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3
```



# burn 

```
./build/index-cli burn 123 --endpoint http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3
```
# balance 

```
./build/index-cli balance index1l97kg5xvpxm0qvhyy6vuej2llg7ydlcp86euda6ewk3uuq4jks0q2rm5qx  --endpoint http://127.0.0.1:9650/ext/bc/2AaCeDDkUdPzqzsKm6iEjiVFWZFVX64izXKdgKrJnJ2FKqWpe3
```



## 


150,000

