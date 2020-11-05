FROM ubuntu:latest

COPY hello-world /usr/bin

WORKDIR /usr/bin

CMD ["hello-world"]

