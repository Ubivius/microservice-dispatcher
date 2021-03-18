#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/player -XPOST -d '{"id":1, "ip":"0.0.0.0"}'
