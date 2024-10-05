FROM ubuntu:latest
LABEL authors="erora"

ENTRYPOINT ["top", "-b"]