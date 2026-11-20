FROM golang

COPY . .
RUN go build .
RUN cp check-test /check-test
CMD /check-test
