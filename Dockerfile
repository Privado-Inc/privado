FROM alpine:latest
ARG VERSION=1.0.0
WORKDIR /home
RUN mkdir -p /home/rules
COPY ./rules ./rules
RUN echo $VERSION >> ./rules/version.txt
