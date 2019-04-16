# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/sbouaazzi/book

# Build the book command inside the container
# Fetch and manage dependencies manually with 'RUN go get' commands
RUN go get github.com/asaskevich/govalidator
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go install github.com/sbouaazzi/book

# Run the book command by default when the container starts
ENTRYPOINT /go/bin/book

# Document that the service listens on port 8080
EXPOSE 8080