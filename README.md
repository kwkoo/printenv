# OpenShift Go multi-stage build example

This example shows you how to perform a multi-stage Go build in OpenShift.

These instructions were adapted from <https://blog.openshift.com/chaining-builds/>.

## Install Golang imagestream

````
oc login -u system:admin

oc create \
  -f https://raw.githubusercontent.com/sclorg/golang-container/master/imagestreams/golang-centos7.json \
  -n openshift
````

## Build

1. Create the builder to build the Go application:

    ````
    oc new-build \
      golang~http://qnap.home:3000/kwkoo/printenv.git \
      --name=builder \
      --context-dir=src \
      -e IMPORT_URL=. \
      -e INSTALL_URL='github.com/kwkoo/printenv'
    ````
2. Create a secondary builder assemble the final image:

    ````
    oc new-build \
      --name=runtime \
      --source-image=builder \
      --source-image-path=/opt/app-root/gobinary:. \
      --dockerfile=$'FROM scratch\nCOPY gobinary /\nUSER 1001\nEXPOSE     8080\nENTRYPOINT ["/gobinary"]' \
      --strategy=docker
    ````
3. Create the application:

    ````
    oc new-app \
      runtime \
      --name=my-application
    ````

Note: Because we used `scratch` as a base image in step 2, we had to include a custom S2I `assemble` script with the source in order to statically link the executable.