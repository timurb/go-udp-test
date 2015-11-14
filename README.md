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

real	0m0.824s
user	0m0.459s
sys	0m0.361s

Running test2: iteration with UDP sending to localhost and no receiver

real	0m3.976s
user	0m1.041s
sys	0m2.866s

Running test3: iteration with UDP sending to localhost and receiver listening for UDP traffic

real	0m6.948s
user	0m1.519s
sys	0m4.382s
./run.sh: line 25: 81051 Terminated: 15          $NC -ul 1234 > /dev/null
```
