#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

SEELIE_DIR="$(cd "$SCRIPT_DIR"/../../ >/dev/null; pwd -P)"
ROOT_DIR="$(cd "$SCRIPT_DIR"/../../../ >/dev/null; pwd -P)"

#echo project root: $ROOT_DIR
#echo seelie dir: $SEELIE_DIR

cd $SEELIE_DIR

if [[ ! -f "$ROOT_DIR/.git/hooks/pre-commit" ]]; then
  mkdir -p $ROOT_DIR/.git/hooks
  cp -f ./script/git/pre-commit-hook.sh $ROOT_DIR/.git/hooks/pre-commit
  chmod +x  $ROOT_DIR/.git/hooks/pre-commit
#  echo "1"
else
  chmod +x  $ROOT_DIR/.git/hooks/pre-commit
#  echo "2"
fi