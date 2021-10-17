#!/bin/bash

dir=`pwd`

build() {
  i=0
  for arg in "$@"
  do
	  pushd $dir/$arg >/dev/null
    make build
    chmod a+x $arg
	  popd >/dev/null
	  let i+=1
  done
  echo total build count: $i
}

# 开始 build 指定的服务
build  advert comment gateway news user