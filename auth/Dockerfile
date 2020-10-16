FROM golang:1.15-alpine AS build_base

WORKDIR /tmp/go-app

COPY . .

RUN go build -o out/auth

RUN ls /tmp/go-app

FROM alpine:3.12.0


WORKDIR /bin

COPY --from=build_base /tmp/go-app/out/auth .

RUN ls
CMD [ "./auth" ]
