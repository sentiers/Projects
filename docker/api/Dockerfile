# Pull base image.
FROM golang:1.16-alpine

# Copy files
COPY . /API_project

# Define working directory.
WORKDIR /API_project

# RUN go build
RUN go build -o api-server

# # Port
# EXPOSE 8080

# # Define default command.
# CMD [ "./api-server" ]

# ====== Docker command ======
# docker build -t apiserver .
# docker images
# docker run --network="host" api

# docker compose