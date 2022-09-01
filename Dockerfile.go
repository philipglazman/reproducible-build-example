FROM debian:bullseye-slim@sha256:139a42fa3bde3e5bad6ae912aaaf2103565558a7a73afe6ce6ceed6e46a6e519

RUN apt update && apt install -y unzip wget tar ca-certificates
RUN wget https://storage.googleapis.com/golang/go1.17.linux-amd64.tar.gz
RUN tar xvf go1.17.linux-amd64.tar.gz
RUN mkdir -p /opt/go && cp -r go /opt/go/go1.17
RUN apt install -y gcc
RUN cd /opt/go/go1.17/src && GOROOT_BOOTSTRAP=/go ./make.bash

ENV PATH "/opt/go/go1.17/bin:$PATH"

ENTRYPOINT ["go"]