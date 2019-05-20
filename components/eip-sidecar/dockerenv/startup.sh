#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

echo "Starting VPN client..."
vpnclient start

echo "Starting EIP-sidecar..."
/nalej/eip-sidecar $@

echo "Waiting"
sleep 6000