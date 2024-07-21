#!/bin/bash

SRC=$1

if ls "$SRC"/tests/test_*.py 1> /dev/null 2>&1; then
  echo "Test found"
  exit 0
  else
  echo "No tests found"
  exit 1
fi
