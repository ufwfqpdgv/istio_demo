FROM alpine

COPY smh-ticket-server /server/
# COPY recommend/config/* /server/config/

WORKDIR /server/

ENTRYPOINT /server/smh-ticket-server
