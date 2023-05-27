FROM oven/bun:0.6.4 AS builder
WORKDIR /app
COPY worker/package.json ./worker/
COPY worker/bun.lockb ./worker/
COPY worker/*.ts ./worker/
COPY GNUmakefile .
RUN apt update
RUN apt install build-essential -y
RUN ls -l .
RUN make setup/worker
RUN make build/worker

FROM oven/bun:0.6.4
WORKDIR /app
ENV SERVER_HOSTNAME=${SERVER_HOSTNAME:-localhost}
ENV SERVER_PORT=${SERVER_PORT:-3000}
COPY --from=builder /app/worker/worker .
CMD ["worker"]
