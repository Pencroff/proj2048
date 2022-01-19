#ls -la
#nodemon --verbose --delay 500ms --watch ".\proj2048\**\*.go" --signal SIGTERM --exec "go run" ".\proj2048\main.go"
# Works below
nodemon --verbose --delay 300ms --watch ".\proj2048\**\*.go" --exec "go run" ".\proj2048\main.go"