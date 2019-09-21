FROM golang

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root
ENV DEBIAN_FRONTEND 'noninteractive'

RUN apt-get update -y
RUN apt-get install -y --no-install-recommends apt-utils

WORKDIR /HighloadHW2

COPY . .

EXPOSE 9090

RUN go get -u

CMD go run main.go