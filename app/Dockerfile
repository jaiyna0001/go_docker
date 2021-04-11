FROM golang:latest

RUN mkdir /go/src/work
# RUN go get -u github.com/oxequa/realize && \
#     go get -u github.com/gorilla/mux
WORKDIR /go/src/work
ADD . /go/src/work

CMD ["go", "run", "app/main.go"]