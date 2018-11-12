#!/bin/sh

set -x

echo "SHTTPD"
wrk -t2 -c32 -d15 -s paths.lua http://127.0.0.1:8080/index.html

echo "CADDY"
wrk -t2 -c32 -d15 -s paths.lua http://127.0.0.1:2015/index.html
