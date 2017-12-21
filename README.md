# openfaas-identicon
> OpenFaaS function to generation Identicons

This is a simple [FaaS](https://github.com/openfaas/faas) function inspired by Bart Fokker's [post](https://blog.bartfokker.nl/identicon/).

## Prerequisite
* Make sure you have deployed a FaaS stack to your cluster using the instructions on the [FaaS repo](https://github.com/alexellis/faas).
* Install [faas-cli](https://github.com/alexellis/faas-cli/) using `brew install faas-cli` or `curl -sSL https://get.openfaas.com | sudo sh`.

## Usage
**Create a new function**
```
$ faas-cli new --lang go identicon
```

**Clone this repo as the function**
```
$ rm -rf identicon

$ git clone git@github.com:ganesshkumar/openfaas-identicon.git identicon
```

**Installing the dependencies**
```
$ cd identicon

$ dep ensure

# As of the build stage will throw an error when the vendor code(dependencies) doesn't comply gofmt
# So let's make sure our dependencies passes the gofmt check

$ gofmt -w vendor/*

$ cd ..
```

**Build and deploy**
```
$ faas-cli build -f identicon.yml

$ faas-cli deploy -f identicon.yml
```

**Testing the function**
```
$ curl localhost:8080/function/identicon --data "openfaas" > sample/openfaas.png
```

![Sample image generated for openfaas](sample/openfaas.png)
