#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.17-alpine AS build_stage
ENV PACKAGE_PATH=/app
RUN mkdir -p $PACKAGE_PATH
RUN mkdir -p $PACKAGE_PATH/configs
WORKDIR $PACKAGE_PATH
ENV CGO_ENABLED=0
ADD go.mod go.sum $PACKAGE_PATH
RUN go mod download
COPY . $PACKAGE_PATH/
RUN cd cmd/telegram-bot && go build
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM alpine:latest AS final_stage
ENV TZ="Asia/Almaty"
ENV PACKAGE_PATH=/app
RUN mkdir -p $PACKAGE_PATH/configs
COPY --from=build_stage $PACKAGE_PATH/cmd/telegram-bot/telegram-bot $PACKAGE_PATH/
COPY --from=build_stage $PACKAGE_PATH/configs/config_telegram_bot.yml $PACKAGE_PATH/configs
WORKDIR $PACKAGE_PATH/
CMD ./telegram-bot
EXPOSE 8080