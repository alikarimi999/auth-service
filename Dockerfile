FROM golang:alpine3.16

RUN apk add gcompat

COPY . /app/
WORKDIR /app
# run the app
# RUN go mod tidy; go build -o application
RUN apk add gcompat; chmod +x auth_service
CMD ["./auth_service" ]