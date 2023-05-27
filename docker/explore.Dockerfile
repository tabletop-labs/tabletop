FROM oven/bun:0.6.4 AS builder
WORKDIR /app
COPY explore/package.json ./explore/
COPY explore/bun.lockb ./explore/
COPY explore/*.ts ./explore/
COPY GNUmakefile .
RUN apt update
RUN apt install build-essential -y
RUN make setup/explore
RUN make build/explore

FROM oven/bun:0.6.4
WORKDIR /app
ENV SERVER_HOSTNAME=${SERVER_HOSTNAME:-localhost}
ENV SERVER_PORT=${SERVER_PORT:-3000}
COPY --from=builder /app/explore/explore .
CMD ["explore"]
