#!/usr/bin/env sh
export GOPATH=/Volumes/B/go_workspace
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN:/usr/local/bin

go run /Volumes/B/go_workspace/src/AmazonMetrics/metrics.go
node /Volumes/B/go_workspace/src/AmazonMetrics/server.js


