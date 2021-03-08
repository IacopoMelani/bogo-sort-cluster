# MARK: build stage

FROM golang:alpine as build

WORKDIR /app/build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

# MARK: final stage

FROM alpine

WORKDIR /app/bogo-sort

COPY --from=build /app/build/bogo-sort-cluster .

CMD [ "./bogo-sort-cluster" ]

EXPOSE 8888