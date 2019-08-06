FROM golang:1.12.7-stretch AS builder

ENV GOBIN /go/bin
ENV GO111MODULE on 

RUN mkdir /app
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/dominik-zeglen
RUN mkdir /go/src/github.com/dominik-zeglen/inkster
ADD . /go/src/github.com/dominik-zeglen/inkster
WORKDIR /go/src/github.com/dominik-zeglen/inkster

RUN make schema

RUN CGO_ENABLED=0 go build -o /app/main manage.go
RUN CGO_ENABLED=0 go build -o /app/migrate migrations/*.go

FROM alpine
WORKDIR /app
RUN mkdir /app/app
COPY --from=builder /app/main /app
COPY --from=builder /app/migrate /app
COPY ./app/graphql.html /app/app
COPY ./config.toml /app

CMD ["/app/main", "runserver"]
