FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /app/krdv
COPY . .
COPY ./frontend /app/krdv/frontend
COPY .env /app/krdv/.env
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
RUN GOOS=linux go build -ldflags="-s -w" -o ./krdv-app ./main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /go/bin
COPY --from=build /app/krdv /go/bin
EXPOSE 4000
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
ENTRYPOINT /go/bin/krdv-app --port 4000