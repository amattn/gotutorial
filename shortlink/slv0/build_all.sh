#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 

COMPONENT="shortlink"
GO_COMMAND="go"

# generate
echo "************************"
echo "$GO_COMMAND generate"
$GO_COMMAND generate

# build on the native or default platform
echo "************************"
echo "building native platform"
$GO_COMMAND build

echo "************************"
echo "vetting native platform"
$GO_COMMAND vet

# test on the native or default platform
echo "************************"
echo "testing native platform"
$GO_COMMAND test

# I like gox as a cross compilation tool: https://github.com/mitchellh/gox
# install with:
# go get github.com/mitchellh/gox

# build for our test or deployment platforms
# normally we do local development on darwin/amd64 and deploy to linux/amd64.
# feel free to add or remove if your situation differs
# also, in the normal case, most of the output of gox is redundant with
# the output from go build above, so in the normal case, we just 
# redirect to a build log

# currently gox doesn't seem to support vgo...

GOX_BUILD_LOG="gox_build.log"
date >> $GOX_BUILD_LOG
echo "************************"
echo "crosscompiling: gox -osarch=\"darwin/amd64\" -osarch=\"linux/amd64\""
if ! gox -osarch="darwin/amd64" -osarch="linux/amd64" >> $GOX_BUILD_LOG 2>&1; then
    echo "FAILURE: gox command failed to build for deployment architecture"
    echo "check gox_build.log for more info"
fi

