FROM golang:1.9

WORKDIR $GOPATH/src/endpoint
COPY ./src/endpoint .

RUN go-wrapper download   
RUN go-wrapper install  

CMD ["go-wrapper", "run"] # ["endpoint"]