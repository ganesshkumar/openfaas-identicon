# openfaas-identicon
> OpenFaaS function to generation Identicons

This is a simple [FaaS](https://github.com/openfaas/faas) function inspired by Bart Fokker's [post](https://blog.bartfokker.nl/identicon/).

## Prerequisite
* Make sure you have deployed a FaaS stack to your cluster using the instructions on the [FaaS repo](https://github.com/alexellis/faas).
* Install [faas-cli](https://github.com/alexellis/faas-cli/) using `brew install faas-cli` or `curl -sSL https://get.openfaas.com | sudo sh`.

## Usage
**Deploy the function**
```
$ faas-cli deploy -image=ganesshkumar/faas-identicon -name=identicon -fprocess="/usr/bin/identicon"
No existing function to remove
Deployed.
URL: http://localhost:8080/function/identicon

200 OK
```

**Testing the function**
```
$ curl localhost:8080/function/identicon --data "openfaas" > sample/openfaas.png
```

![Sample image generated for openfaas](sample/openfaas.png)
