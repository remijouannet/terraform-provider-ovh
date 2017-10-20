Terraform Provider for OVh DNS (unofficial)
==================

Requirements
------------

-   [Terraform](https://www.terraform.io/downloads.html) 0.10.2 
-   [Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Install
---------------------

Download the binary and put it in the same folder than terraform binary

```
$ wget https://github.com/remijouannet/terraform-provider-ovh/releases/download/v0.1/terraform-provider-ovh_darwin_amd64_v0.1.zip
$ unzip terraform-provider-ovh_darwin_amd64_v0.1.zip
$ mv terraform-provider-ovh_darwin_amd64_v0.1/terraform-provider-ovh-v0.1 ~/bin/
$ chmod +x ~/bin/terraform-provider-ovh-v0.1
```

add the following to ~/.terraformrc

```
providers {
  ovh = "/home/remi/bin/terraform-provider-ovh-v0.1"
}
```


Build without docker
---------------------

Clone repository to: `$GOPATH/src/github.com/remijouannet/terraform-provider-ovh`

```
$ mkdir -p $GOPATH/src/github.com/remijouannet; cd $GOPATH/src/github.com/remijouannet
$ git clone git@github.com:remijouannet/terraform-provider-ovh
```

Enter the provider directory and build the provider

```
$ cd $GOPATH/src/github.com/remijouannet/terraform-provider-ovh
$ make build
```

Build with docker
---------------------

build the docker image

```
$ make docker-image
```

build the binaries, you'll find all the binaries in pkg/

```
$ make docker-build
```


How to use
---------------------


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
* ovh.tf example

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


How to Test
---------------------

create a script name env.sh, it will contain the following

```
export TF_ACC="yes"
export OVH_APPLICATION_KEY="arqzergfazfeef"
export OVH_APPLICATION_SECRET="aertedytfghftyhytghbgfv"
export OVH_CONSUMER_KEY="tyfghjhuyghvbjhuytghbjygh"
export OVH_ZONE="testdemo.com"
```

then use the makefile : make test
