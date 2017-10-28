FROM golang:1.9.2

RUN mkdir -p /go/src/github.com/Adaptech/gameoflife
WORKDIR /go/src/github.com/Adaptech/gameoflife
COPY . /go/src/github.com/Adaptech/gameoflife
RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install
ENV PORT 3001
EXPOSE 3000
CMD gin run
