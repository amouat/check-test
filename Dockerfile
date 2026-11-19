FROM golang

COPY . .
RUN go build .
CMD check-test
