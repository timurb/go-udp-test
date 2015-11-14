# Test on how much CPU consumes UDP send to localhost

## Running

* Make sure you have Golang installed.
* Run `./build.sh` to build code
* Run `./run.sh` to run tests

## Results

Here are the results obtain on my MacBook for 1000000 iterations:

```
$ ./run.sh
Running test1: iteration

real	0m0.111s
user	0m0.104s
sys	0m0.006s

Running test2: iteration with UDP sending to localhost and no receiver

real	0m3.236s
user	0m0.495s
sys	0m2.662s

Running test3: iteration with UDP sending to localhost and receiver listening for UDP traffic

real	0m3.802s
user	0m0.573s
sys	0m2.907s
./run.sh: line 25: 84353 Terminated: 15          $NC -ul 1234 > /dev/null
```
