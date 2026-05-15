FROM swift:5.9-jammy AS build

WORKDIR /build

COPY Package.swift ./
RUN swift package resolve

COPY Sources ./Sources
RUN swift build -c release --static-swift-stdlib

FROM ubuntu:jammy

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

RUN useradd --create-home --system vapor

USER vapor

WORKDIR /app

COPY --from=build --chown=vapor:vapor /build/.build/release/App ./

EXPOSE 8080

ENTRYPOINT ["./App"]
CMD ["serve", "--env", "production", "--hostname", "0.0.0.0", "--port", "8080"]
