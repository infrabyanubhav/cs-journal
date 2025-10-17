#!/bin/bash
/Users/anubhavsharma/Desktop/packet-sniffer/myenv/bin/python3 --version >/dev/null 2>&1 || {
  echo "Python venv not found at myenv. Adjust path in sniff.sh."
  exit 1
}

sudo /Users/anubhavsharma/Desktop/packet-sniffer/myenv/bin/python3 /Users/anubhavsharma/Desktop/packet-sniffer/sniffer/main.py \
  > /Users/anubhavsharma/Desktop/packet-sniffer/sniffer/output.txt 2>&1 &

echo "Sniffer running (PID $!) — logs: sniffer/output.txt"