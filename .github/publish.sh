#!/usr/bin/env bash

set -xe

export GOARCH="amd64"
export GOOS="darwin"
go build -o committer.amd64 committer.go

export GOARCH="arm64"
go build -o committer.arm64 committer.go

lipo committer.amd64 committer.arm64 -create -output committer
