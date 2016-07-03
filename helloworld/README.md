This sample is forked from [https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/helloworld](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/helloworld).

## Running on [PCF Dev](https://docs.pivotal.io/pcf-dev)

``` console
$ cf push
```

App runs on [http://hello-go-helloworld.local.pcfdev.io](http://hello-go-helloworld.local.pcfdev.io)

## Running locally

```
$ export VCAP_APPLICATION={}
$ export VCAP_SERVICES={}
$ go run helloworld.go
```
App runs on [http://localhost:4000](http://localhost:4000)
