#!/bin/bash

# Determine script directory
DIR="$(cd "$(dirname "$0")" && pwd)"

# Prefer project venv's Python, fallback to system python3
PY="$DIR/../myenv/bin/python3"
if [ ! -x "$PY" ]; then
  PY="$(command -v python3)"
fi

echo "Sniffer running"
# Run the sniffer from this script's directory, write output next to it
"$PY" "$DIR/main.py" > "$DIR/output.txt" 2>&1 &