FROM ubuntu:latest
# FROM debian:bullseye-slim

ARG LANG=jpn

RUN apt update \
    && apt install -y \
    ca-certificates \
    libtesseract-dev \
    tesseract-ocr

RUN apt install -y software-properties-common \
    && add-apt-repository ppa:longsleep/golang-backports \ 
    && apt install -y golang

ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin
ENV PORT=8989

COPY . $GOPATH/src/github.com/sgash708/OCR_Go
WORKDIR $GOPATH/src/github.com/sgash708/OCR_Go

# RUN go get -v ./... && go install .

RUN apt install -y tesseract-ocr-$LANG
# CMD ["OCR_Go"]