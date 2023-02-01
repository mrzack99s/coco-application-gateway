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
RUN mkdir -p /app/logs

COPY --from=build /build/certs /app/certs

COPY --from=build /build/conf /app/conf

COPY --from=build /build/coco-application-gateway /app/coco-application-gateway

COPY --from=build /build/coraza.conf /app/coraza.conf

COPY --from=build /build/owasp /app/owasp

WORKDIR /app

RUN chmod +x coco-application-gateway

ENTRYPOINT [ "./coco-application-gateway", "run" ]
