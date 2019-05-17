FROM golang:latest

RUN mkdir -p $GOPATH/src/golang.org/x/ \
    && cd $GOPATH/src/golang.org/x/ \
    && git clone https://github.com/golang/net.git \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/op/go-logging

WORKDIR /ntc/src

CMD ["go", "build", "-o", "/ntc/bin/NDP-task-courier.b", "./"]
