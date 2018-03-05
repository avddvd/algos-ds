#!/bin/bash -xe
DIR="github.com/avddvd/algos-ds"
docker build -t prep .
docker run -ti \
  -v `pwd`/src:/go/src/${DIR}/ \
  prep bash
