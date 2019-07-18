FROM golang:1.12-alpine as build
RUN apk add --update --no-cache git
WORKDIR /work
ADD go.mod go.sum ./
WORKDIR /work/cmd
ADD cmd/go.mod cmd/go.sum ./
RUN go mod download
WORKDIR /work
ADD *.go ./
ADD proto proto
ADD cmd cmd
WORKDIR /work/cmd
RUN go build -o /ubivolt

FROM alpine
LABEL maintainer="Adam Engebretson <adam@enge.me>"
RUN apk add --update --no-cache ca-certificates
COPY --from=build /ubivolt /ubivolt
EXPOSE 5000
VOLUME /data
ENTRYPOINT ["/ubivolt", "serve"]
CMD ["/data/ubivolt.db"]