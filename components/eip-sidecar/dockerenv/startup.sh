#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

echo "Starting VPN client..."
vpnclient start

echo "Starting EIP-sidecar..."
/nalej/eip-sidecar $@

echo "Enabling IP forwarding..."
echo 'net.ipv4.ip_forward=1' >> /etc/sysctl.conf
#sysctl -p

while true; do
    echo "alive"
    sleep 100000
done
