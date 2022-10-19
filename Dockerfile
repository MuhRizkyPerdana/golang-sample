FROM golang:alpine3.16 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /studidevops

FROM golang:alpine3.16
WORKDIR /
COPY --from=build /studidevops /studidevops
EXPOSE 80
ENTRYPOINT ["/studidevops"]
