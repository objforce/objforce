

### salesforce 多租户设计
参考: https://developer.salesforce.com/wiki/multi_tenant_architecture

data sparce table 采用 MySQL? hbase ?
mt_indexes, mt_unique_indexes 采用 MySQL?


salesforce 采用了 golang & GRPC
参考这篇文章: https://www.cncf.io/case-studies/salesforce/

关键问题总结:
我们过去在基于JSON的集成中遇到的一个痛点是，它们需要在每一方进行大量的协商，并且对于向后不兼容的更改来说可能很脆弱.