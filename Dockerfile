# Sttart from base image 1.12.13:
FROM golang:1.13.7

ENV LOG_LEVEL=info

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/vadym-vasilyev/collector-go

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/github.com/vadym-vasilyev/collector-go

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o collector-sercice .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./collector-sercice"]