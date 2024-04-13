go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go build -o hello
./hello