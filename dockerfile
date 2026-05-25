FROM golang:1.25.6

# set the working directory
WORKDIR /go/src/app

# Copy the dependency files first
COPY go.mod go.sum ./
RUN go mod download

# copy the source code
COPY . .

# EXPOSE the port
EXPOSE 8000

# build the go app
RUN go build -o main cmd/main.go

# Run the executable
CMD [ "./main" ]