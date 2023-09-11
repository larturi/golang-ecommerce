git add .
git commit -m "Commit"
git push

GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main  

# chmod +x deploy.sh
# ./deploy.sh
