#Build stage
FROM golang:1.19-alpine as build

COPY . /build

WORKDIR /build

RUN go mod tidy

RUN go install main.go

RUN go build -o coco-application-gateway main.go

# Prod stage
FROM alpine:3 as prod

RUN mkdir -p /app/certs

COPY --from=build /build/certs /app/certs

COPY --from=build /build/conf/config.yaml /app/conf/config.yaml

COPY --from=build /build/conf/routing.yaml /app/conf/routing.yaml

COPY --from=build /build/coco-application-gateway /app/coco-application-gateway

WORKDIR /app

RUN chmod +x coco-application-gateway

ENTRYPOINT [ "./coco-application-gateway", "run" ]
