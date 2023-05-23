# grpc-helloworld-demo

**接口测试**

```shell
curl --location 'http://localhost:8090/v1/arithmetic/add' \
--header 'accept: application/json' \
--header 'Content-Type: application/json' \
--data '{  "a": "15",  "b": "17"}'
```