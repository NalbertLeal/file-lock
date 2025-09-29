mkdir ./bin
OOS=windows GOARCH=amd64 go build -o ./bin/file-lock_win_amd64.exe *.go
OOS=windows GOARCH=386 go build -o ./bin/file-lock_win_386.exe *.go
OOS=linux GOARCH=amd64 go build -o ./bin/file-lock_linux_amd64 *.go
OOS=linux GOARCH=386 go build -o ./bin/file-lock_linux_386 *.go