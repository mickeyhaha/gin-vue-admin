export GO111MODULE=off
export GOPATH=/opt/go
export PATH=/opt/go/bin:$PATH
export GOPROXY=https://goproxy.io
export CGO_ENABLED=0   
export GOOS=windows  
export GOARCH=386

go build main.go 


