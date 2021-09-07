FROM ubuntu:latest

ENV LANG=jpn

RUN apt update \
    && apt install -y \
    ca-certificates \
    libtesseract-dev \
    tesseract-ocr \
    tesseract-ocr-$LANG \
    software-properties-common

RUN add-apt-repository ppa:longsleep/golang-backports \
    && apt install -y golang

ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /var/www/html