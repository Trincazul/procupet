FROM caddy
WORKDIR /usr/src/app
COPY go.mod go.sum main.go ./
RUN go mod download 
COPY . .
CMD ["go", "run", "main.go"]
EXPOSE 9000


