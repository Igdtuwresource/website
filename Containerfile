FROM golang:1.18 as build

WORKDIR /root

COPY . .

RUN curl -fsSL -o ./caddy 'https://caddyserver.com/api/download?os=linux&arch=amd64' && \
    chmod +x ./caddy

RUN go run main.go

### 

FROM debian:11

RUN apt-get update && apt-get install -y \
      libnss3-tools && \
    rm -rf /var/cache/apt/*

RUN useradd --create-home website
USER website
WORKDIR /home/website

COPY --from=build --chown=website:website /root/site /home/website/site
COPY --from=build --chown=website:website /root/caddy /home/website/caddy
COPY Caddyfile .

RUN ./caddy validate

EXPOSE 2019

CMD ["/home/website/caddy", "run"]
