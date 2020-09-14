FROM golang:latest as build-env
WORKDIR /go/src/ml_daily_record
ADD . /go/src/ml_daily_record
RUN go build -mod=vendor -o /go/app

FROM golang:latest as prod-env
WORKDIR /go/src/ml_daily_record
COPY --from=build-env /go/src/ml_daily_record/resources/ resources
COPY --from=build-env /go/src/ml_daily_record/config-dev.yml config-dev.yml
COPY --from=build-env /go/app .
EXPOSE 8080
CMD ["./app"]
