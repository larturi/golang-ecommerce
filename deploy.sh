git add .
git commit -m "Commit"
git push

GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
zip lambda-handler.zip bootstrap  

# chmod +x deploy.sh
# ./deploy.sh
