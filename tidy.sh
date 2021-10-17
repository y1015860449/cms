#!/bin/bash

dir=`pwd`

test() {
  i=0
  for arg in "$@"
  do
	  pushd $dir/$arg >/dev/null
	  go mod tidy
	  popd >/dev/null
	  let i+=1
  done
  echo total tidy count: $i
}

test advert comment gateway news user

