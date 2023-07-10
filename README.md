# Gopan-基于go-zero的分布式网盘项目

## User服务



## Upload服务

### 普通上传接口

- 可以选择本地或者minio存储或者腾讯云COS进行存储

![image-20230708153557534](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230708153557534.png)

- 使用kafka异步处理Mysql存储文件元信息

异步处理，可以让用户不用等待文件元信息存入mysql的所需的时间，提高用户体验。针对高并发场景，提高吞吐量。

- 使用批量消息聚合提升kafka性能

![image-20230708153904061](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/image-20230708153904061.png)

之前每向kafka发送一条消息就会产生一次网络 IO 和一次磁盘 IO，做消息聚合后，比如聚合 100 条消息后再发送给 Kafka，这个时候 100 条消息才会产生一次网络 IO 和磁盘 IO，这样大大提高 Kafka 的吞吐和性能。并且有聚合时间兜底，就算消息数量达不到聚合要求，超过聚合最大时间也会聚合当前所有消息发送给Kafka

### 秒传

根据文件的sha1值判断，先从file表里面查出是否有该sha1值，如果有就只需要将file表里的file添加到对应userfile表里面就可以了，这样就实现了秒传的功能