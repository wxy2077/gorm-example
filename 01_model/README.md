
## model层 对应表字段

只写纯SQL操作，尽可能的封装更多筛选条件的基础操作，来减少代码冗余操作，如`user`表的`First`方法，封装很多过滤的参数，方便更多业务场景重复使用。