FROM golang:1.8

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install app/price_list
RUN go install app/pricing
RUN go install app/ranking

EXPOSE 6667:6667

CMD ["./run.sh"]
