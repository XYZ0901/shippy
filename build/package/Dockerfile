FROM golang:1.15.3

RUN mkdir /shippy

WORKDIR /shippy

COPY ../.. .

RUN go build -o shippy-vessel ./vessel-service

CMD ["./shippy-vessel"]