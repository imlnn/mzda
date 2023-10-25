FROM alpine:3 AS build

RUN apk add --no-cache go
RUN go version

WORKDIR /app
COPY . .
#RUN "go mod download && go build -o /mzda cmd/mzda/main.go && chmod +x /mzda"
#RUN "go mod tidy && go mod vendor"
#RUN "go build -o /mzda cmd/mzda/main.go"
#RUN "chmod +x /mzda"

CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"

#FROM alpine:3

#COPY --from=build /mzda /mzda

#CMD ["./mzda"]
