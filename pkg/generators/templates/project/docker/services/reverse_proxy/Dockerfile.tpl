FROM nginx:1.19.7-alpine
# default environment configuration that can be overridden at run time
ENV MINIO_EXTERNAL_HOST=${MINIO_EXTERNAL_HOST:-minio.localhost}
ENV MINIO_INTERNAL_HOST=${MINIO_INTERNAL_HOST:-minio}
ENV MINIO_HTTP_PORT=${MINIO_HTTP_PORT:-9001}
ENV REDPANDA_CONSOLE_EXTERNAL_HOST=${REDPANDA_CONSOLE_EXTERNAL_HOST:-redpanda.localhost}
ENV REDPANDA_CONSOLE_INTERNAL_HOST=${REDPANDA_CONSOLE_INTERNAL_HOST:-redpanda-console}
ENV REDPANDA_CONSOLE_HTTP_PORT=${REDPANDA_CONSOLE_HTTP_PORT:-8080}
ENV PROMETHEUS_EXTERNAL_HOST=${PROMETHEUS_EXTERNAL_HOST:-prometheus.localhost}
ENV PROMETHEUS_INTERNAL_HOST=${PROMETHEUS_INTERNAL_HOST:-prometheus}
ENV PROMETHEUS_HTTP_PORT=${PROMETHEUS_HTTP_PORT:-9090}
ENV GRAFANA_EXTERNAL_HOST=${GRAFANA_EXTERNAL_HOST:-grafana.localhost}
ENV GRAFANA_INTERNAL_HOST=${GRAFANA_INTERNAL_HOST:-grafana}
ENV GRAFANA_HTTP_PORT=${GRAFANA_HTTP_PORT:-3000}
ENV REVERSE_PROXY_HTTP_PORT=${REVERSE_PROXY_HTTP_PORT:-8080}
ENV REVERSE_PROXY_DNS_RESOLVER_ADDRESS=${REVERSE_PROXY_DNS_RESOLVER_ADDRESS:-127.0.0.11}

COPY ./templates/nginx.conf /etc/nginx/templates/nginx.conf

RUN apk add --update apache2-utils

CMD sh -c 'envsubst "\
      \${CONTROL_UI_EXTERNAL_HOST} \
      \${CONTROL_UI_INTERNAL_HOST} \
      \${CONTROL_UI_HTTP_PORT} \
      \${CONTROL_API_EXTERNAL_HOST} \
      \${CONTROL_API_INTERNAL_HOST} \
      \${CONTROL_API_HTTP_PORT} \
      \${MINIO_EXTERNAL_HOST} \
      \${MINIO_INTERNAL_HOST} \
      \${MINIO_HTTP_PORT} \
      \${REDPANDA_CONSOLE_EXTERNAL_HOST} \
      \${REDPANDA_CONSOLE_INTERNAL_HOST} \
      \${REDPANDA_CONSOLE_HTTP_PORT} \
      \${PROMETHEUS_EXTERNAL_HOST} \
      \${PROMETHEUS_INTERNAL_HOST} \
      \${PROMETHEUS_HTTP_PORT} \
      \${GRAFANA_EXTERNAL_HOST} \
      \${GRAFANA_INTERNAL_HOST} \
      \${GRAFANA_HTTP_PORT} \
      \${REVERSE_PROXY_HTTP_PORT} \
      \${REVERSE_PROXY_DNS_RESOLVER_ADDRESS}" \
    < /etc/nginx/templates/nginx.conf \
    > /etc/nginx/nginx.conf \
    && nginx -g "daemon off;"'
