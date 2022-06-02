FROM debian:11

RUN apt-get update && apt-get install -y \
      curl \
      libnss3-tools && \
    rm -rf /var/cache/apt/*

RUN useradd --create-home website
USER website
WORKDIR /home/website

RUN curl -fsSL \
      -o ./hugo.tar.gz \
      'https://github.com/gohugoio/hugo/releases/download/v0.100.0/hugo_extended_0.100.0_Linux-64bit.tar.gz' && \
    tar -xzf ./hugo.tar.gz && \
    rm ./hugo.tar.gz

RUN curl -fsSL -o ./caddy 'https://caddyserver.com/api/download?os=linux&arch=amd64' && \
    chmod +x ./caddy

COPY --chown=website:website config.toml .
COPY --chown=website:website content ./content
# COPY --chown=website:website data ./data
COPY --chown=website:website layouts ./layouts
# COPY --chown=website:website resources ./resources
COPY --chown=website:website scripts ./scripts
COPY --chown=website:website static ./static
COPY --chown=website:website Caddyfile .

RUN ./hugo
RUN ./caddy validate

EXPOSE 2015

CMD ["/home/website/caddy", "run"]
