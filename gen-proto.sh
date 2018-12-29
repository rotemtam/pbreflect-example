#!/bin/bash

protoc --go_out=$GOPATH/src/ *.proto
