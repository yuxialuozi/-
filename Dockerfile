# 使用 OpenTelemetry Collector 官方镜像作为基础镜像
FROM otel/opentelemetry-collector-contrib:latest

# 将本地配置文件复制到容器的配置目录中
COPY ./otel.yaml /etc/otel-collector-config.yaml

# 设置默认命令以启动 OpenTelemetry Collector
CMD [ "--config=/etc/otel-collector-config.yaml" ]
