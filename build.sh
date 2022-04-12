echo "Step 1/3: Building application"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

echo "Step 2/3: Building Docker image"

docker-compose build go

echo "Step 3/3: Running Docker container"

docker-compose up -d