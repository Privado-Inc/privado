FROM alpine:latest
WORKDIR /home
RUN mkdir -p /home/rules
COPY ./rules ./rules
