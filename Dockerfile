FROM debian:bullseye-slim as build

WORKDIR /build
LABEL org.opencontainers.image.source="https://github.com/simbafs/coscup-attendance"
LABEL org.opencontainers.image.authors="Simba Fs <me@simbafs.cc>"

ENV PATH="/usr/local/go/bin:/usr/local/node-v18.17.1-linux-x64/bin:$PATH"

# install dependencies
RUN apt-get update && \
    apt-get install -y xz-utils wget ca-certificates --no-install-recommends && \
    # config ssl
    mkdir -p /etc/ssl/certs && \
    update-ca-certificates --fresh && \
    # install go v1.22.4
    wget -q https://go.dev/dl/go1.22.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz 

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main . && \
    mkdir -p /build/data


FROM scratch 

WORKDIR /app/data

COPY --from=build /build/main /app/main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/data /app/data

EXPOSE 3000
CMD [ "/app/main" ]
