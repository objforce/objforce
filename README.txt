# 元模型引擎



### Best Practice
golang 微服务 最佳实践

解决核心架构问题:
* 容器化, docker, k8s
* 分布式配置中心集中管理, 目前基于 apollo config扩展

代码规范性问题
* 严格唯一性编码规范, 代码框架模版化
* 采用 uber 的轻量级di框架，保证代码可测试性



## salesforce ID 格式
62进制格式
前6位是prefix
其中前3位表示 objType的id, 总共有 62^3 = 238328种表达
四到六位的6F0为Org Id的第4到6位
第四位代表 orgId 所属的 instance

第6-15共9位连续自增, 那么这里需要一个类似snowflake的分布式连续自增生成器，只需要针对租户本身自增


关于ID生成
https://zhuanlan.zhihu.com/p/107470205