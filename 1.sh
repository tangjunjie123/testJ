kill -9 $(lsof -i:7777 -t)
go build -o hello
BUILD_ID=DONTKILLME
nohup ./hello &>hello.log &
