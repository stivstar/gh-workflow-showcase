# syntax=docker/dockerfile:1
# Alpine is chosen for its small footprint
FROM golang:alpine

RUN mkdir /dockapp
WORKDIR /dockapp

# Copy and download necessary Go modules
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Build the golang app and expose port to outside world
RUN go build -o /golang-garrulous-gopher
EXPOSE 1323

CMD [ "/golang-garrulous-gopher" ]