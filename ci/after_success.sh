#!/bin/bash

if [ "$RUN_TYPE" = "coverage" ]; then
   coveralls-lcov --repo-token "$COVERALLS_REPO_TOKEN" --service-name travis-pro ./build/merged.info

   if [ "$TRAVIS_BRANCH" = "master" ] && [ "$TRAVIS_PULL_REQUEST" = "false"]; then
      COMMIT_HASH=`git rev-parse --short HEAD`

      # Clone koinos-types-golang, add changes, commit, and push
      cd ~
      git clone git@github.com:koinos/koinos-types-golang
      cd koinos-types-golang

      cp $TRAVIS_BUILD_DIR/build/generated/golang/src/github.com/koinos/koinos-types-golang/* ./
      cp $TRAVIS_BUILD_DIR/tests/golang/* ./

      if ! git diff --exit-code; then
         git add -u
         git commit -m "Update for koinos-types commit $COMMIT_HASH"
         git push
      fi
   fi
fi
