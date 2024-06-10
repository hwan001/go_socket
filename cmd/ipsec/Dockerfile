FROM alpine:latest

RUN apk add --no-cache strongswan iproute2 bash

COPY strongswan.conf /etc/strongswan/strongswan.conf
COPY ipsec.conf /etc/ipsec.conf
COPY ipsec.secrets /etc/ipsec.secrets
COPY serverKey.pem /etc/ipsec.d/private/serverKey.pem

CMD ["bash", "-c", "ipsec start --nofork"]
