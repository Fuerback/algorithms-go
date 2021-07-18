FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go mod download

ENV PORT :8000

RUN go build

ENTRYPOINT ["go-quicksort"]
CMD ["./go-quicksort"]