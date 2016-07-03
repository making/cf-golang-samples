This sample is forked from [https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/static_files](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/static_files).

## Running on [PCF Dev](https://docs.pivotal.io/pcf-dev)

``` console
$ cf push
```

App runs on [http://hello-go-staticfiles.local.pcfdev.io](http://hello-go-staticfiles.local.pcfdev.io)

## Running locally

```
$ export VCAP_APPLICATION={}
$ export VCAP_SERVICES={}
$ go run staticfiles.go
```

App runs on [http://localhost:4000](http://localhost:4000)
