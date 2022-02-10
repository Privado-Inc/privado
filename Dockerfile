# Container image that runs your code
FROM docker.elastic.co/beats-dev/golang-crossbuild:1.17.1-darwin

COPY . /app/

ENTRYPOINT ["build/build.sh"]
