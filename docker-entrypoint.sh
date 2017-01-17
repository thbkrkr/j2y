#!/bin/bash -eu

case $1 in

  diff)

    diff <(cat $1 | json2yaml -mode y2j | jq -M .) <(cat $2 | json2yaml -mode y2j | jq -M .)
  ;;

  *)
    exec json2yaml
  ;;

esac
