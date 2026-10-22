CGO_ENABLED=1 GOOS=linux go build -o nawst
docker build -t nawst .