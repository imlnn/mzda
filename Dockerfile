FROM alpine:3 AS build

RUN apk add --no-cache go
RUN go version

WORKDIR /app
COPY . .
RUN go mod vendor && go mod download && go build -o /mzda cmd/mzda/main.go && chmod +x /mzda

FROM alpine:3
COPY --from=build /mzda /mzda

CMD ["./mzda"]
