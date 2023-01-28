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
RUN cd service-store/cmd/ && go build -o main
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM alpine:latest AS final_stage
ENV TZ="Asia/Almaty"
RUN mkdir -p /service-store/config
RUN mkdir -p /service-store/migrations
COPY --from=build_stage app/service-store/cmd/ service-store/
COPY --from=build_stage app/service-store/config/config.yml service-store/config
COPY --from=build_stage app/service-store/migrations/ service-store/migrations/
CMD service-store/main
EXPOSE 8002