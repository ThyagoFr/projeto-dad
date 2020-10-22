FROM golang:alpine
ENV JWT_KEY=91e0dcf69ddfcbc436c6f4c9fec29ee5
ENV AWS_ACCESS_KEY=""
ENV AWS_SECRET_ACCESS_KEY=""
ENV AWS_REGION=us-east-1
ENV AWS_BUCKET_NAME_BOOK=booksdad2020
ENV AWS_BUCKET_NAME_READER=readersdad2020
ENV AWS_QUEUE=emaildad
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN ls
RUN go build src/main.go
EXPOSE 8090
CMD ["/app/main"]