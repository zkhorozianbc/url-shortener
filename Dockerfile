FROM golang:1.16


WORKDIR /app

COPY . .

RUN go mod download

RUN go install 

EXPOSE 9808

CMD ["url-shortener"]
