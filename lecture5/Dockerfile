FROM golang
WORKDIR /app
COPY . .
RUN go get
RUN go build -o api .
CMD ["./api"]