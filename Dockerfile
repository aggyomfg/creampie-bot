# syntax=docker/dockerfile:experimental
FROM golang:1.13 AS modules
WORKDIR $GOPATH/src/github.com/aggyomfg/creampie-bot
COPY go.mod .
COPY go.sum .
RUN --mount=type=cache,target=/root/.cache/go-build GO111MODULE=on go mod download

FROM golang:1.13 as build
COPY --from=modules $GOCACHE $GOCACHE
COPY --from=modules $GOPATH/pkg/mod $GOPATH/pkg/mod
WORKDIR $GOPATH/src/github.com/aggyomfg/creampie-bot
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o=/bin/bot ./cmd/bot

FROM alpine as image
COPY --from=build /bin/bot /bin/bot
### Временно конфиг
ADD ./.env /.env
EXPOSE 8080
CMD ["/bin/bot"]