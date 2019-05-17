#
# Copyright (C) 2019 Nalej Group - All Rights Reserved
#

#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

echo "Starting VPN client..."
vpnclient start

echo "Starting EIP-sidecar..."
/nalej/eip-sidecar $@
