#!/bin/bash
isExistApp=`pgrep hello`
if [[ -n  $isExistApp ]]; then
  pgrep hello | xargs -I {} kill {}
fi
