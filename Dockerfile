FROM golang:alpine3.16

COPY . /app/
WORKDIR /app
# run the app
RUN go mod tidy; go build -o application
RUN chmod +x application
CMD [ "./application" ]