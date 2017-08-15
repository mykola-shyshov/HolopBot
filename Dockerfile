FROM golang:1.8

WORKDIR /go/src/app
COPY . /go/src/app

RUN go-wrapper download holop  # "go get -d -v ./..."
RUN go-wrapper install  holop  # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]
