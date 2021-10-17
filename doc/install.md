#### 安装go-micro v2
1、 安装micro工具包
* go get github.com/micro/micro/v2  
* 执行micro help看到如下信息表示已经安装成功

2、自动生成代码  
* cd go-workspace
* micro new --gopath=false hello

3、安装protobuf
* go get -u github.com/golang/protobuf/proto
* go get -u github.com/golang/protobuf/protoc-gen-go
* go get github.com/micro/micro/v2/cmd/protoc-gen-micro

4、proto文件编译
* protoc --proto_path=. --micro_out=. --go_out=. proto/hello/hello.proto