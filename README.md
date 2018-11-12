# Static HTTPD
*Stefan Arentz, December 2018*

This is a small web server that loads all content in memory and serves without touching the disk.

Results below are from a 2018 MacBook Air. This is probably not a good benchmark setup.

Caddy as a (Go) reference:

```
+ wrk -t2 -c32 -d15 -s paths.lua http://127.0.0.1:2015/index.html
Running 15s test @ http://127.0.0.1:2015/index.html
  2 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.72ms    2.77ms  37.98ms   72.99%
    Req/Sec     2.06k   173.96     2.47k    69.05%
  61640 requests in 15.06s, 24.30GB read
Requests/sec:   4091.71
Transfer/sec:      1.61GB
```

Here is `shttpd`, running as `go run -root testdata`:

```
Running 15s test @ http://127.0.0.1:8080/index.html
  2 threads and 32 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.96ms    1.63ms  15.90ms   69.76%
    Req/Sec     2.56k   258.62     3.28k    66.55%
  76514 requests in 15.01s, 31.14GB read
Requests/sec:   5098.00
Transfer/sec:      2.08GB
```
