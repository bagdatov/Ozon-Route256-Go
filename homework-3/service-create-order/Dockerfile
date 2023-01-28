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
RUN cd service-create-order/cmd/ && go build -o main
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM alpine:latest AS final_stage
ENV TZ="Asia/Almaty"
RUN mkdir -p /service-create-order/config
RUN mkdir -p /service-create-order/migrations
COPY --from=build_stage app/service-create-order/cmd/ service-create-order/
COPY --from=build_stage app/service-create-order/config/config.yml service-create-order/config
COPY --from=build_stage app/service-create-order/migrations/ service-create-order/migrations/
CMD service-create-order/main
EXPOSE 8000