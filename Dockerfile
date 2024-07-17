FROM golang
WORKDIR /home/server
COPY . .
EXPOSE 4000
CMD ["go","run","main.go"]