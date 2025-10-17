## Packet Sniffer (ARP Scan)

A minimal ARP network scanner using Scapy that discovers devices on a subnet and writes results to `sniffer/output.txt`.

### Prerequisites
- macOS
- Python 3.13 (already in `myenv/` virtualenv)
- Admin privileges (ARP scanning typically requires `sudo`)

### Setup
1. (Optional) Activate the virtual environment:
```bash
source myenv/bin/activate
```
2. Install dependencies:
```bash
# Using requirements.txt (recommended)
pip install -r sniffer/requirements.txt

# Or install manually
pip install scapy
```

### Usage
- Run via script (recommended) from the project root:
```bash
chmod +x sniffer/sniff.sh
sudo sniffer/sniff.sh
```
This will run the scanner in the background and write results to `sniffer/output.txt`.

- Run directly with Python from the project root:
```bash
sudo myenv/bin/python sniffer/main.py > sniffer/output.txt 2>&1
```

### Changing the subnet
`sniffer/main.py` scans `192.168.1.0/24` by default. To scan another subnet, edit the default:
```python
results = arp_scan(network="10.0.0.0/24")
```

### Output
Example lines in `sniffer/output.txt`:
```text
IP: 192.168.1.1, MAC: 6c:08:31:07:b7:ef
IP: 192.168.1.35, MAC: 5e:d2:26:23:79:45
```

### Troubleshooting
- If you see permission errors, ensure you used `sudo`.
- If no hosts appear, verify you selected the correct subnet and your interface is up.
- On macOS, you may be prompted to allow terminal network permissions.

### Notes
- `.gitignore` excludes the virtualenv and `sniffer/output.txt` from version control.
- The script auto-detects the venv Python at `myenv/bin/python3` if available.
- If the venv is missing, it falls back to your system `python3`.
