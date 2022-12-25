FROM golang:1.18.3-alpine3.15 as build
#LABEL builder=true
RUN mkdir -p /go/src/github.com/bsromr/modern-webapp-w-devops
WORKDIR /go/src/github.com/bsromr/modern-webapp-w-devops
COPY . .
RUN ls -lah
RUN pwd
RUN go mod download
RUN go build -o app cmd/web/*.go

##ONLY 1 stage for now
# FROM alpine:edge 
# #LABEL builder=false
# WORKDIR /app
# RUN mkdir templates
# COPY --from=build /go/src/github.com/bsromr/modern-webapp-w-devops .
# RUN cp ./app cmd/web/
# EXPOSE 8080
# WORKDIR /app/cmd/web
# RUN ["./app"]