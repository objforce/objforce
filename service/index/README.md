

### salesforce 多租户设计
参考: https://developer.salesforce.com/wiki/multi_tenant_architecture


MT_Objects表：用于存储对象的定义，包括了其唯一的ID、从属的组织的ID、名字等。在Salesforce系统中的所有“租户”组织的对象定义同时保存在这一个表中

MT_Fields表：用于存储对象字段的信息，包括了其唯一的ID、从属的组织的ID、从属的对象的ID、名字、字段类型等。在Salesforce系统中的所有“租户”组织包含的字段同时保存在这一个表中

MT_Data表：用于存储各个对象的实际数据，相当于传统关系数据库中的一行记录。在Salesforce系统中所有“租户”组织中的数据同时保存在这一个表中

MT_Clobs表：CLOB是“character large objects”的简称。这些对象可以存储长达32000字符的数据。当MT_Data表中存在的数据包含过多的字符时，系统会将这些长字符数据存储在MT_Clobs表中，而在MT_Data表中存储一个MT_Clobs表某一行的ID值，从而精简MT_Data表

MT_Indexes表：这是一个数据透视表（Pivot table），用于存储MT_Data表中数据的索引。该表主要用于提高数据的搜索效率

MT_Unique_Indexes表：该表和MT_Indexes表类似，也是一个数据透视表，也存储了数据的索引。不同之处在于该表中的数据索引保持着唯一性。在用户给某字段增加唯一约束时，系统会将该字段的内容记录在此表中。当用户插入重复的值的时候，系统会根据该表的内容给出警告信息

MT_Fallback_Indexes表：MT_Fallback_Indexes表中保存了所有数据的名字（Name）。在特殊情况下，Salesforce有可能无法完成用户需要的搜索。在这种情况下，Salesforce会启用备用搜索机制，直接查询MT_Fallback_Indexes表中的数据，从而给出名字符合搜索条件的数据



随着云计算的发展及应用软件的成熟, 软件即服务 (Software as a Service, SaaS)[1]作为云计算的一种应
用形式越来越受到重视. 多租户数据架构是搭建 SaaS应用平台的关键技术之一, 不仅需要在数据库层面实现租户之间数据的隔离[2], 还要满足租户的定制需求.
目前几种典型的多租户共享存储模型, 包括透视表、稀疏表、块表及块折叠, 都能保障租户数据的隔离性和可定制性的需求, 但仍存在各自的不足. 例如,


参考:
https://github.com/rahgadda/Salesforce/blob/bffb4702b45056a05ad5a759bd1d81d5b0dddfad/02-Modules/01-Force.com/01-MultiTenant.md