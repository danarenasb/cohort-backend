# stage needs the Go image to build the binary
FROM golang:1.17-alpine AS stage 
WORKDIR /app
COPY go.* *.go ./
RUN go mod download
RUN GOARCH=amd64 go build -o wylee-backend

FROM gcr.io/distroless/base
COPY --from=stage wylee-backend wylee-backend
EXPOSE 8080
CMD [ "wylee-backend" ]