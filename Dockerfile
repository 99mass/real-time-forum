#Use base image with golang latest version
FROM golang:latest

LABEL VERSION="latest"
LABEL DESCRIPTION="FORUM APP"
LABEL AUTHORS="@osamb - @mouhametadiouf - @dalassan - @alo - @ssambadi"

#Define the working directory
WORKDIR /forumApp

#Copy the go.mod and go.sum files into the container
COPY go.mod .
COPY go.sum .

#Download go dependencies
RUN go mod download

#Copy source code to container
COPY . .

#Compile the application
RUN go build -o forum

#Expose application listening port
EXPOSE 8080

#Application start command
CMD ["./forum"]