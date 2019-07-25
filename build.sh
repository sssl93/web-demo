#!/bin/bash

tag='v1.0.0'

docker build -f ./Dockerfile -t web-demo:$tag .