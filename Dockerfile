FROM golang:1

WORKDIR /go/src/github.com/hjkelly/budget-category-service

EXPOSE 8080

RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -vendor-only

COPY server.go .
COPY common common
COPY categories categories
COPY views views

RUN go install

CMD ["budget-category-service"]
