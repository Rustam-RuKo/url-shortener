FROM golang:alpine as builder

RUN mkdir / builde/

ADD . /build/

RUN go build -o main .


FROM alpine

RUN adduser  -S -D -H / app appuser

USER appuser

COPY . app/

COPY --from=builder / build / main/ app/

WORKDIR / app

EXPOSE 3000
CMD ["./main"]