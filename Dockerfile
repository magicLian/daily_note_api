FROM golang:1.13.5-alpine3.10 as build-env
WORKDIR /go/src/ml_daily_record
ADD . /go/src/ml_daily_record
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -s' -o /go/app

FROM alpine:3.11.5 as prod-env
WORKDIR /go/src/ml_daily_record
COPY --from=build-env /go/src/ml_daily_record/resources/ resources
COPY --from=build-env /go/src/ml_daily_record/config-dev.yml config-dev.yml
COPY --from=build-env /go/app .
EXPOSE 8080
CMD ["./app"]
