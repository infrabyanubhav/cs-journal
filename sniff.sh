#!/bin/bash

echo "Sniffer running"
# Run the sniffer
sudo python sniffer/main.py > sniffer/output.txt 2>&1 &