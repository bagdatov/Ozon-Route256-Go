#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.17-alpine AS build_stage
ENV PACKAGE_PATH=/app
RUN mkdir -p $PACKAGE_PATH
WORKDIR $PACKAGE_PATH
ENV CGO_ENABLED=0
ADD go.mod go.sum $PACKAGE_PATH
RUN go mod download
COPY . $PACKAGE_PATH/
RUN cd service-monitoring-order/cmd/ && go build -o main
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM alpine:latest AS final_stage
ENV TZ="Asia/Almaty"
RUN mkdir -p /service-monitoring-order/config
RUN mkdir -p /service-monitoring-order/migrations
COPY --from=build_stage app/service-monitoring-order/cmd/ service-monitoring-order/
COPY --from=build_stage app/service-monitoring-order/config/config.yml service-monitoring-order/config
COPY --from=build_stage app/service-monitoring-order/migrations/ service-monitoring-order/migrations/
CMD service-monitoring-order/main
EXPOSE 8001