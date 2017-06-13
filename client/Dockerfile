FROM scratch



MAINTAINER Ton Schoots <ton@maiastra.com>

EXPOSE 8080

ENV  STARGRAPH_SERVER_IP localhost

ADD ./dist /dist
ADD ./html /html
ADD ./sk_stargraph_client /sk_stargraph_client

CMD ["./sk_stargraph_client"]