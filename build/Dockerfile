FROM golang:1.18.3-alpine3.15 as build
#LABEL builder=true
RUN mkdir -p /go/src/github.com/bsromr/modern-webapp-w-devops
WORKDIR /go/src/github.com/bsromr/modern-webapp-w-devops
COPY . .
RUN ls -lah
RUN pwd
RUN go mod download
RUN go build -o app cmd/web/*.go
EXPOSE 8080
CMD ["./app"]

# # FROM alpine:edge 
# FROM golang:1.18.3-alpine3.15
# # #LABEL builder=false
# WORKDIR /app
# COPY --from=build /go/src/github.com/bsromr/modern-webapp-w-devops/app .
# EXPOSE 8080
# CMD ["./app"]