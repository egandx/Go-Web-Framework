# gRPC

## 一、Preparations
```
1、Protocol Buffers: 
    1.1 github地址: https://github.com/protocolbuffers/protobuf
    1.2 golang库: https://github.com/golang/protobuf
    1.3 gotutorial官方文档: https://developers.google.com/protocol-buffers/docs/gotutorial
    
2、Mac 安装包安装protoc(https://segmentfault.com/a/1190000039732564)

3、开始尝试跑Demo。相关步骤和错误处理参考Grpc初尝试.pdf

4、grpc-go：https://github.com/grpc/grpc-go

protoc --go-grpc_out=./ --go-grpc_opt=paths=source_relative ProductInfo.proto

```
    
## 二、Overview

### 1、总体过程
```
1.1 客户端发送数据（以字节流的方式）

1.2 服务端接收，并解析。根据约定明确要执行什么操作，完成后把结果返回给客户端。
```
    
### 2、RPC作用
```
2.1 将上述的过程进行封装，使其操作更加的优化

2.2 使用一些大家都熟悉认可的协议，使其更加的规范化

2.3 做成框架，直接或间接的产生利益
```


## 三、Start

### 1、创建proto文件
```protobuf
syntax = "proto3";
package services;
option go_package="../services";

message ProdRequest {
  int32 prod_id = 1;  //商品ID
}

message ProdResponse {
  int32 prod_stock = 1; //商品库存
}

service ProdService {
  rpc GetProdStock (ProdRequest) returns (ProdResponse);
}
```

```shell
protoc --go-grpc_out=../services --go-grpc_opt=paths=source_relative Prod.proto
```

### 2、创建gRPCServer和gRPCClient

### 3、自签证书
```shell
3.1 执行OpenSSL
    >openssl

3.2 生成grpcserver.key 私钥文件
    OpenSSL>genrsa -des3 -out grpcserver.key 2048

3.3 创建证书请求,生成grpcserver.csr
    OpenSSL>req -new -key grpcserver.key -out grpcserver.csr

3.4 删除密码
    OpenSSL>rsa -in grpcserver.key -out grpcserver_no_password.key

3.5 生成grpcserver.crt,自签证书完成
    OpenSSL>x509 -req -days 365 -in grpcserver.csr -signkey grpcserver_no_password.key -out grpcserver.crt

```


