#
# Copyright (C) 2018 Nalej Group - All Rights Reserved
#

# iron/go is the alpine image with only ca-certificates added
FROM iron/go

# Now just add the binary\
RUN mkdir /nalej
EXPOSE 5544/tcp

# Now just add the binary
COPY edge-inventory-proxy /nalej/

ENTRYPOINT ["./nalej/edge-inventory-proxy"]