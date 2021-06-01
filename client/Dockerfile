FROM golang:1.16

WORKDIR /client

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 5000 to the outside world
EXPOSE 5000

RUN go build .

CMD ["./client"]