#!/usr/bin/env bash
set -euo pipefail

echo "GoGate bootstrap status"
echo "======================"

for p in docs management apps core engines infra .github; do
  if [ -e "$p" ]; then
    echo "[ok] $p"
  else
    echo "[missing] $p"
  fi
done
