FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go bulid -o main main.go

EXPOSE 8080
CMD ["/app/main"]