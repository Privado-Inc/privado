FROM alpine:latest
ARG VERSION=1.0.0
WORKDIR /home
RUN mkdir -p /home/rules
COPY . /home/rules
RUN echo $VERSION >> /home/rules/version.txt
