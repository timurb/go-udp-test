#!/bin/sh

if which nc>/dev/null; then
  NC=nc
elif which netcat>/dev/null; then
  NC=netcat
else
  echo "No netcat/nc command available"
  exit 1
fi

echo "Running test1: iteration"
time ./test1

echo
echo "Running test2: iteration with UDP sending to localhost and no receiver"
time ./test2


echo
echo "Running test3: iteration with UDP sending to localhost and receiver listening for UDP traffic"
$NC -ul 1234 > /dev/null &
time ./test2

killall $NC
