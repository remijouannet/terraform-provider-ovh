# terraform_osc

## howto
if you want to use the OVH API you have to generate:
    * An Application Key
    * An Application Secret Key
    * A Consumer Key

you can generate all three on this page 
https://eu.api.ovh.com/createToken/

the following rights are needed for this plugin
```
GET    : /domain/zone/*
PUT    : /domain/zone/*
POST   : /domain/zone/*
DELETE : /domain/zone/*
```
* tf example

```
provider "ovh" {
    application_key = "azrzqrgqvvdsfgsfffgc"
    application_secret = "aztfqsqfsdcsdqezrfdvcx"
    consumer_key = "aergfvdsrgtfbvcretfgd"
}

resource "ovh_domain_zone_record" "test" {
    zone = "testdemo.ovh"
    subDomain = "test"
    fieldType = "A"
    ttl = "3600"
    target = "0.0.0.0"
}
```


## MakeFile
* build : make build

* install : make install

* testing : create a script name env.sh, it will contain the following

```
export TF_ACC="yes"
export OVH_APPLICATION_KEY="arqzergfazfeef"
export OVH_APPLICATION_SECRET="aertedytfghftyhytghbgfv"
export OVH_CONSUMER_KEY="tyfghjhuyghvbjhuytghbjygh"
export OVH_ZONE="testdemo.com"
```

then use the makefile : make test
