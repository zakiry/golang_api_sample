# golang_api_sample

sample request
```shell
$ curl -X POST -H "Content-Type: application/json" -d '{"id": 2}' "localhost:3000/twirp/sample_proto.UserService/GetUser"
{"id":"2","name":"piyo piyo","age":30}
```