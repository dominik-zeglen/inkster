FROM golang:1.10.3-stretch AS app-builder

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/dominik-zeglen
RUN mkdir /go/src/github.com/dominik-zeglen/inkster
ADD . /go/src/github.com/dominik-zeglen/inkster
WORKDIR /go/src/github.com/dominik-zeglen/inkster

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -o /app/main app/main.go


FROM node:8.11.4-alpine AS ui-builder

RUN mkdir /app
run mkdir /src
RUN mkdir /src/app
ADD ./panel /src/app
WORKDIR /src/app

RUN npm install
RUN npm run build

FROM alpine
WORKDIR /app
COPY --from=app-builder /app/main /app
COPY --from=ui-builder /src/app/build /app/panel/build

CMD ["/app/main"]
