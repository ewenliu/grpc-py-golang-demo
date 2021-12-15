## pb使用说明

golang
- 安装protoc，参考https://www.cnblogs.com/baoshu/p/13488106.html
- 安装protoc-gen-go，go get google.golang.org/protobuf
- 安装grpc go getgoogle.golang.org/grpc
- 编译proto文件 protoc --go_out=plugins=grpc:. Product.proto

py
- 安装grpcio-tools， pip install grpcio-tools
- 编译proto文件 python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. Product.proto

跨语言调用需要注意的地方
- proto文件中的变量名建议一律使用驼峰，如果使用下划线可能会有跨语言调不通的现象