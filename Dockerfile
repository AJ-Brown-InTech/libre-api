# Specify the base image for the go app.
FROM --platform=linux/amd64 golang:1.18
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/AJ-Brown-InTech/libre-api/
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq
# used for live reloading changes
RUN go install github.com/cosmtrek/air@latest
#install all that other ish (packages)
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

# Compile the binary exe for our app.
RUN go build -o main .
# Start the application.
#CMD ["./main"]
CMD ["air", "-c", ".air.toml"]