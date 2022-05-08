FROM golang:1.17-alpine3.15 as build-stage
WORKDIR /go/src/app
COPY . .
RUN go get -v
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"'

FROM scratch as production-stage
COPY --from=build-stage /go/src/app/server /app/
WORKDIR /app
EXPOSE 3000
CMD ["./server"]
