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


# WORKDIR /go/src/app
# RUN mkdir -p /work/xprincipia/backend
# WORKDIR go/src/work/xprincipia/backend

# Add backend folder
ADD . /go/src/work/xprincipia/backend

# Install the Backend
RUN go install work/xprincipia/backend


# # Set the entry point of the container to the application executable
ENTRYPOINT /go/bin/backend

# Expose the application on port 10000.
# This should be the same as in the app.conf file
EXPOSE 10000



