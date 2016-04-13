FROM alpine:3.3

LABEL tthumb_version="R1"

RUN apk add --update openssh-client imagemagick

EXPOSE 80

COPY ./tthumb /tthumb
COPY ./presets.json /presets.json

WORKDIR /

CMD ["/tthumb", "-path", "/media", "-port", ":80"]

