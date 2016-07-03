This sample is forked from [https://github.com/GoogleCloudPlatform/golang-samples/blob/master/docs/managed_vms/memcache](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/docs/managed_vms/memcache).

## Running on [PCF Dev](https://docs.pivotal.io/pcf-dev)

``` console
$ cf create-service p-redis shared-vm redis-db
$ cf push
```

App runs on [http://hello-go-redis.local.pcfdev.io](http://hello-go-redis.local.pcfdev.io)

## Running on [Pivotal Web Services](https://run.pivotal.io)

``` console
$ cf create-service rediscloud 30mb redis-db
$ cf push --random-route
```
App runs on [http://hello-go-redis-{random-word}.cfapps.io](http://hello-go-redis-{random-word}.cfapps.io)

## Running locally

```
$ export VCAP_APPLICATION={}
$ export VCAP_SERVICES='{"p-redis":[{"credentials":{"hostname":"localhost","port":6379,"password":""},"name":"redis-db"}]}'
$ go run redis.go
```

App runs on [http://localhost:4000](http://localhost:4000)
