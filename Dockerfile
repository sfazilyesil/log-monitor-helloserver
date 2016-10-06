FROM alpine

ADD helloserver /

ENTRYPOINT ["/helloserver"]