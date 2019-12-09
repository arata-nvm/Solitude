#!/bin/bash

OPT=$*
TARGET=bin/solitude

try() {
  expected="$1"
  input="$2"

  echo "$input" | $TARGET $OPT > tmp.ll
  lli tmp.ll
  actual=$?

  if [ "$actual" == "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$input => $expected expected, but got $actual"
    exit 1
  fi
}

try 0 "0"
try 42 "42"

try 4 "2+2"
try 5 "8-3"
try 7 "7-0"
