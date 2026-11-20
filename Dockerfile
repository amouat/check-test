FROM golang

COPY . .
RUN go build .
RUN pwd
CMD ["/go/check-test"]
