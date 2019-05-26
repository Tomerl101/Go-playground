# We specify the base image we need for our
# go application
FROM golang:1.12.0-alpine3.9
# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD web /app
# We specify that we now wish to execute
# any further commands inside our /app
# directory
WORKDIR /app
# we run go build to compile the binary
# executable of our Go program (create main file from all files in the web dir)
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]

### BUILD THE DOCKER IMAGE ###
#$ docker build -t my-go-app . (from root dir)

### RUN THE DOCKER IMAGE ###
# $ docker run -p 8080:3000 -it my-go-app
# then go the localhost:8080 it will redirect to port 3000 on the container where the web server port is running

# $docker ps
# $docker kill <container_id>