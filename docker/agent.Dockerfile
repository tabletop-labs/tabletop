FROM oven/bun:0.6.4 AS builder
WORKDIR /app
COPY agent/package.json ./agent/
COPY agent/bun.lockb ./agent/
COPY agent/*.ts ./agent/
COPY GNUmakefile .
RUN apt update
RUN apt install build-essential -y
RUN make setup/agent
RUN make build/agent

FROM oven/bun:0.6.4
WORKDIR /app
ENV SERVER_HOSTNAME=${SERVER_HOSTNAME:-localhost}
ENV SERVER_PORT=${SERVER_PORT:-3000}
COPY --from=builder /app/agent/agent .
CMD ["agent"]
