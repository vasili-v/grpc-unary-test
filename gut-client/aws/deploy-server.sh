#!/bin/bash

wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
sudo tar -xzf /home/ec2-user/go1.8.3.linux-amd64.tar.gz -C /usr/local/
sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
go version

sudo yum install -y git
mkdir -p src/github.com/vasili-v
cd src/github.com/vasili-v/
git clone https://github.com/vasili-v/grpc-unary-test
cd grpc-unary-test/gut-server

export GOPATH=$HOME
export PATH=$GOPATH/bin:$PATH
go get -u
go install
