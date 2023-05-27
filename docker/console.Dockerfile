FROM oven/bun:0.6.4 AS builder
WORKDIR /app
COPY console/package.json ./console/
COPY console/bun.lockb ./console/
COPY console/*.ts ./console/
COPY GNUmakefile .
RUN apt update
RUN apt install build-essential -y
RUN make setup/console
RUN make build/console

FROM oven/bun:0.6.4
WORKDIR /app
ENV SERVER_HOSTNAME=${SERVER_HOSTNAME:-localhost}
ENV SERVER_PORT=${SERVER_PORT:-3000}
COPY --from=builder /app/console/console .
CMD ["console"]
