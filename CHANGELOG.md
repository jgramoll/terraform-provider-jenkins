# terraform-provider-jenkins Change Log

## 0.2.0

[e62a9a3] File path (#2)

## 0.1.1

[9083e51] make sure to return nil instead of error on read
[352126c] make sure provider doesn't fail on read if property or action doesn't exit
[5139baf] add local dev link info to readme
[0515f40] update readme to have clearer install vs build
[3747ae0] update readme dependencies to have clearer gomod command
[b079ff7] Update readme with more about credentials and envvars
[b6b812f] update importer usage to build using go get
[1cc8c5f] clean up readme todo

## 0.1.0

[10fa096] add import script to importer to get tf state (#1)
[e9d0d95] update importer to output .tf code files for job
[d05d39a] refactor remaining client unmarshal switches
[0b28faf] remove importer binary
[c1ef766] add ensure id functionality to importer
[e80be59] add jira project property
[d88ee11] initial importer
[76854a7] allow for resource import
[bd0c1d3] refactor type checks on dynamics to better handle errors
[fb503f0] refactor unmarshal to use map instead of switch
[2c46df9] add job_datadog_job_property
[3b009a4] fix job config serialize test
[6f8a8d1] allow for  gerrit trigger plugin attr
[9588a33] fix losing git scm resource
[5b54ec1] add job declarative actions to provider
[995ac14] add job config actions to client

## 0.0.3

[fabf3bd] fix id delimiter to be less command character

## 0.0.2

[1fb1065] fix rest of initial resources should be functional now
[50bf38b] (origin/master) get gerrit trigger events
[885918b] finish gerrit trigger project and branch
[4e2f3e4] fix build discarder strategies and tests
[3450264] fix build discarder property
[f5d9e25] fix client tests
[cf1a857] add more types
[a5d34f1] add git scm

## 0.0.2-alpha

- fix install script

## 0.0.1-alpha

- first version