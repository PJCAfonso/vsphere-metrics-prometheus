#! /usr/bin/env bash

# Copyright (c) 2018 VMware, Inc. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# 	http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#rm -rf ./vendor
#rm glide.lock
rm -f ./vsphere-metrics-prometheus
#glide up

grep -R --exclude-dir vendor --exclude-dir .git --exclude build.sh TODO ./

if [ "$1" == "which" ]; then
    if [ -f Dockerfile_LINUX ]; then
        echo "Building Alpine"
    elif [ -f Dockerfile_ALPINE ]; then
        echo "Building Linux"
    fi
    exit 0
elif [ "$1" == "switch" ]; then
    if [ -f Dockerfile_ALPINE ]; then
        echo "Enabling Alpine"
        mv Dockerfile Dockerfile_LINUX
        mv Dockerfile_ALPINE Dockerfile
    elif [ -f Dockerfile_LINUX ]; then
        echo "Enabling Linux"
        mv Dockerfile Dockerfile_ALPINE
        mv Dockerfile_LINUX Dockerfile
    fi
    exit 0
fi

if [ -f Dockerfile_ALPINE ]; then
    echo "Building Linux"
    GOOS=linux GOARCH=amd64 go build .
elif [ -f Dockerfile_LINUX ]; then
    echo "Building Alpine"
    docker run --rm -e GOPATH=/usr/src/go -v ~/go:/usr/src/go -w /usr/src/go/src/github.com/dvonthenen/vsphere-metrics-prometheus golang:1.10.2-alpine3.7 go build -v
fi

docker build -t dvonthenen/vsphere-metrics-prometheus .
if [ "$1" == "push" ]; then
    docker push dvonthenen/vsphere-metrics-prometheus:latest
fi

exit 0