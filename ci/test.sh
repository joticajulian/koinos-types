#!/bin/bash

set -e
set -x

if [ "$RUN_TYPE" = "test" ]; then
   # C++ tests
   pushd build/tests/cpp
   ctest -j3 --output-on-failure
   popd

   # Golang tests
   go test -v ./tests/golang
   go test -v ./build/...

   # Compare multilingual outputs
   python3 programs/canonical-output/check_canonical_output.py --lang-dir build/programs/canonical-output/lang --test-data programs/koinos-types/json/test_data.json

   golint -set_exit_status ./...
fi
