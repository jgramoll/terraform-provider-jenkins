# terraform-provider-jenkins
Terraform Provider to manage jenkins jobs

## Build and install ##

### Dependencies ###

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

[Go modules](https://github.com/golang/go/wiki/Modules) are used for dependency management.  To install all dependencies run the following:

`export GO111MODULE=on`
`go mod vendor`

### Install ###

You will need to install the binary as a [terraform third party plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).  Terraform will then pick up the binary from the local filesystem when you run `terraform init`.

```sh
curl -s https://raw.githubusercontent.com/jgramoll/terraform-provider-jenkins/master/install.sh | bash
```

## Usage ##

### Credentials ###

$jenkins_url/user/$username/configure

### resources ###

```terraform
provider "jenkins" {
  address = "${var.jenkins_address}"
}

resource "jenkins_job" "premerge" {
  name = "Premerge checks"
}

```
