# About
Each language needs to implement a webserver on port 8080 which has "/get" route for get requests that does 50 binary search for "123"  on a sorted array with 1e6 members.  
Also they need to have a way to benchmark the search function with 1e6 iterations and print the average execute time in nano seconds.

# Technologies
## Go
- Router: [Echo](https://github.com/labstack/echo)  
To be fair I limited cpu cores to 1.
#### How to install?
```
cd go
go mod install
go run main.go // for serving
go test -bench=. // for benchmark
```
How to test with apache benckmark?  
```
ab -n 20000 -c 1000 "http://127.0.0.1:8080/get"
```
## PHP
- Web server: [ReactPHP](https://github.com/reactphp) (We can't implement it with cgi interface)
#### How to install?
```
cd php
composer install
php server.php // for webserver
php benchmark.php // for benchmark
```
How to test with apache benckmark?  
```
ab -n 20000 -c 1000 "http://127.0.0.1:8080/get"
```

## NodeJS
- Web framework: [ExpressJS](https://expressjs.com/) 
#### How to install?
```
cd node
npm install
node main.js // for webserver
node benchmark.js // for benchmark
```
How to test with apache benckmark?  
```
ab -n 20000 -c 1000 "http://127.0.0.1:8080/get"
```
# My Results
## Go 1.5
+ Function Execution time: 100ns
+ Requests per second:    15470.45 [#/sec] (mean)
+ Time per request:       65.106 [ms] (mean)
## PHP 8.0
+ Function Execution time: 3000ns
+ Requests per second:    2572.90 [#/sec] (mean)
+ Time per request:       388.667 [ms] (mean)
## NodeJS 14.0
+ Function Execution time: 160ns
+ Requests per second:    7434.52 [#/sec] (mean)
+ Time per request:       134.508 [ms] (mean)





