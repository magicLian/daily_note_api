#!/usr/bin/env bash

docker build -t harbor.arfa.wise-paas.com/appbuy/appbuy-api:1.4.0.11 .
docker push harbor.arfa.wise-paas.com/appbuy/appbuy-api:1.4.0.11
