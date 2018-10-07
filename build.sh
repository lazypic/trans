# OS별로 빌드함.
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/trans trans.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/trans.exe trans.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/trans trans.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && tar -zcvf ../trans_linux.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../trans_windows.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../trans_darwin.tgz . && cd -
