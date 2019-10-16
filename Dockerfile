FROM golang:1.13.1-alpine3.10

RUN apk --no-cache add alpine-sdk coreutils cmake \
  && adduser -D app \
  && echo "app ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
  && mkdir /app \
  && chown -R app:app /app

COPY . /app

WORKDIR /app

USER app

RUN go build

CMD ["/app/gomsg"]
