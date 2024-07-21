#!/bin/bash

SRC=$1

if ls "$SRC"/tests/test_*.py 1> /dev/null 2>&1; then
  echo "Test found"
  else
  echo "No tests found in $SRC"
fi

exit 0
