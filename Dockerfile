# Start from a base image with Go and Git installed

FROM golang:1.17-alpine3.15

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to the container

COPY go.mod go.sum ./

# Download the dependencies

RUN go mod download

# Copy the rest of the application code to the container

COPY . .

# RUN the application
RUN go build -o main .

#Expose the port that the application listen on
EXPOSE 8085

#Set the entrypoint command for the container
CMD [ "./main" ]