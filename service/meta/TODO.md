

### 问题
1. gorm 类型转换问题
复合对象转json字段

2. 参考 salesforce的架构实现, 初步准备用 mycat or vitess 实现多租户
在微服务系统中，由于微服务众多，目前采用schema隔离.
采用database隔离需要在一个比较内聚的系统比较合适.


数据类型问题
前端传递 json
api 接受 json, protobuf 传递 json 到 service

service 的json数据结合 元数据，得到真实数据类型
根据真实数据类型转型，底层字段一律 []byte 存储，index 服务需要还原类型存储