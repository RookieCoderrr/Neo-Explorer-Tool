FROM golang:1.15.6

ENV GO111MODULE="on"

#这个代理好用，不会用redis报错问题
ENV GOPROXY="https://proxy.golang.org"

ARG RT

RUN echo $RT

ENV RUNTIME=$RT

RUN mkdir application

COPY . ./application

WORKDIR "application"

RUN  go build -o main ./app/ws/src.go

EXPOSE 2026

CMD ["./main"]
