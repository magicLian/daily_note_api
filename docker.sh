#!/usr/bin/env bash

docker build -t registry.cn-hangzhou.aliyuncs.com/magiclian/daily_note:0.0.0.1 .
docker push registry.cn-hangzhou.aliyuncs.com/magiclian/daily_note:0.0.0.1
