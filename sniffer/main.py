from scapy.all import ARP, Ether, srp

def arp_scan(network="192.168.1.0/24"):
    # Create ARP request packet
    arp_request = ARP(pdst=network)
    # Create Ethernet frame
    broadcast = Ether(dst="ff:ff:ff:ff:ff:ff")
    # Combine them
    packet = broadcast / arp_request

    # Send packet and capture the response
    answered, unanswered = srp(packet, timeout=2, verbose=0)

    # Process the answered list
    devices = []
    for sent, received in answered:
        devices.append({'ip': received.psrc, 'mac': received.hwsrc})

    return devices


if __name__ == "__main__":
    results = arp_scan()
    for device in results:
        print(f"IP: {device['ip']}, MAC: {device['mac']}")
