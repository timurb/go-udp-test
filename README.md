# Test on how much CPU consumes UDP send to localhost

## Running

* Make sure you have Golang installed.
* Run `./build.sh` to build code
* Run `./run.sh` to run tests

## Results

Here are the results obtained on my MacBook for 1000000 iterations:

```
$ go version
go version go1.3.3 darwin/amd64
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

Here are the results obtained on Ubuntu 12.04 running on c4.large Amazon instance (go version is different):
```
$ go version
go version go1.5.1 linux/amd64
$ ./run.sh
Running test1: iteration
0.08user 0.00system 0:00.07elapsed 111%CPU (0avgtext+0avgdata 25184maxresident)k
0inputs+0outputs (0major+1636minor)pagefaults 0swaps

Running test2: iteration with UDP sending to localhost and no receiver
1.49user 3.88system 0:03.86elapsed 139%CPU (0avgtext+0avgdata 27760maxresident)k
0inputs+0outputs (0major+1801minor)pagefaults 0swaps

Running test3: iteration with UDP sending to localhost and receiver listening for UDP traffic
1.84user 6.34system 0:08.49elapsed 96%CPU (0avgtext+0avgdata 27600maxresident)k
0inputs+0outputs (0major+1789minor)pagefaults 0swaps
```
