FROM siomiz/softethervpn:4.29-ubuntu

COPY dockerenv/startup.sh /startup.sh

RUN apt-get update && apt-get install -y isc-dhcp-client && rm -rf /var/lib/apt/lists/* && mkdir /config && chmod 0755 /startup.sh

COPY eip-sidecar /nalej/

ENTRYPOINT ["/startup.sh"]
