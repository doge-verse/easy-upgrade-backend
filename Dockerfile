FROM golang:stretch as golang-builder
WORKDIR /go/src/github.com/doge-verse/easy-upgrade-backend
COPY . .
COPY cmd/main.go ./main.go
#RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=golang-builder /go/src/github.com/doge-verse/easy-upgrade-backend/app .
COPY --from=golang-builder /go/src/github.com/doge-verse/easy-upgrade-backend/docs ./docs
COPY --from=golang-builder /go/src/github.com/doge-verse/easy-upgrade-backend/template ./template
EXPOSE 8080
ENTRYPOINT ["./app"]