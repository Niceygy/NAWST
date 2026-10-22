CGO_ENABLED=0 GOOS=linux go build -o nawst
docker build -t nawst .