# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.7.3

# # Create the directory where the application will reside
# RUN mkdir /app

# ADD . /go/src/work/xprincipia/backend
# # Copy the application files (needed for production)




# Install Github libraries
RUN go get github.com/golang/glog
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm 
RUN go get github.com/joho/godotenv 
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/appleboy/gin-jwt.v2
RUN go get gopkg.in/gin-gonic/gin.v1

RUN echo ${GOPATH}
WORKDIR /go/src/app
# RUN mkdir -p /work/xprincipia/backend
# WORKDIR go/src/work/xprincipia/backend

ADD . /go/src/app

# WORKDIR /app

RUN ls

# ADD . /go/src/work/xprincipia/backend
# RUN go install /go/src/app/gin
# RUN go install /go/src/app/gorm


RUN go build ./server.go
# RUN cd /go/bin

# RUN pwd


ENTRYPOINT /go/bin/app

EXPOSE 10000

# CMD ["go", "run" ,"/go/src/work/xprincipia/backend/server.go"]

# ENTRYPOINT /go/bin/backend

# Expose the application on port 8080.
# This should be the same as in the app.conf file
# EXPOSE 10000

# # Set the entry point of the container to the application executable
# ENTRYPOINT /go/bin/basic

# CMD go run /backend/server.go -a | xp
