FROM ubuntu:latest

ARG LANG=jpn

RUN apt update \
    && apt install -y \
    ca-certificates \
    libtesseract-dev \
    tesseract-ocr

RUN apt install -y software-properties-common \
    && add-apt-repository ppa:longsleep/golang-backports

RUN apt install -y golang

ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /var/www/html

RUN apt install -y tesseract-ocr-$LANG