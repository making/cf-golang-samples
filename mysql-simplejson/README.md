This sample is forked from [https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/cloudsql](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/docs/managed_vms/cloudsql).

## Running on [PCF Dev](https://docs.pivotal.io/pcf-dev)

``` console
$ cf create-service p-mysql 512mb mysql-db
$ cf push
```

App runs on [http://hello-go-mysql.local.pcfdev.io](http://hello-go-mysql.local.pcfdev.io)

## Running on [Pivotal Web Services](https://run.pivotal.io)

[This line](https://github.com/making/cf-golang-samples/blob/master/mysql-simplejson/mysql.go#L60) needs to be changed to `cleardb`...

``` console
$ cf create-service cleardb spark mysql-db
$ cf push --random-route
```
App runs on [http://hello-go-mysql-{random-word}.cfapps.io](http://hello-go-mysql-{random-word}.cfapps.io)

## Running locally

```
$ export VCAP_APPLICATION={}
$ export VCAP_SERVICES='{"p-mysql":[{"credentials":{"uri":"mysql://root:@localhost:3306/demo"},"name":"mysql-db"}]}'
$ go run mysql.go
```

App runs on [http://localhost:4000](http://localhost:4000)
