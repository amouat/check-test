FROM golang

COPY . .
RUN go build .
CMD ["/go/check-test"]
