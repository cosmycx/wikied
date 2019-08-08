FROM  golang as builder

WORKDIR /wikied

ENV GOPATH /go
ENV GO111MODULE on

COPY go.mod . 
COPY go.sum . 
RUN go mod download
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM scratch 

COPY --from=builder /wikied/app /wikied/app

EXPOSE 4444 
ENTRYPOINT [ "/wikied/app" ]
