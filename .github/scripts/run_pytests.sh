#!/bin/bash

SRC=$1

if ls "$SRC"/tests/test_*.py 1> /dev/null 2>&1; then
  echo "Test found"
  pushd $SRC

  python -m pip install --upgrade pip
  pip install pytest
  if [ -f requirements.txt ]; then pip install -r requirements.txt; fi

  python -m pytest --verbose
  popd
else
    echo "No tests found in $SRC"
    exit 0
fi


