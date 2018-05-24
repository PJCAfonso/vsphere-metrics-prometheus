FROM alpine:3.7
MAINTAINER David vonThenen <davidvonthenen@gmail.com>

ADD vsphere-metrics-prometheus /
EXPOSE 9444

ENTRYPOINT ["/vsphere-metrics-prometheus"]
