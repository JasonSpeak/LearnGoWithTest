### SAP面试总结

1. redis缓存穿透的解决方案
    - 为热点数据加锁
2. O(nlogn)的排序算法
    - 快排、堆排(均不稳定)
    - 归并排序(稳定)
3. golang中sort使用了哪种排序方法，仅仅是快排吗？
    - 使用了堆排序、希尔排序、插入排序
4. 快排的实现原理
    - 
5. go解决并发冲突的方案
    - channel
    - 加锁
6. 关闭一个channel后，对其读写会怎样
    - 关闭的channel仍然可以读取数据，但是当其中的数据读取完之后，继续读的话会读取到一堆默认值
    - 读取channel时最好同时获取两个返回值，第二个返回值是`bool`值，可以反映该channel是否已关闭
7. Tcp协议中timewait的作用
    - timewait是指在Tcp四次挥手中主动关闭连接的一方在发送完最后一次挥手之后，主动关闭连接的一方所处的状态
    - timewait状态的持续时间为2MSL，MSL是“报文最大生存时间”，可为30s,1min或2min,2MSL就是两倍的这个时间
    - timewait的作用：
        - 保证客户端发送的最后一个挥手到达服务器，如果没到达，服务端就会重发第三次挥手
        - 保证本次连接的所有报文段从网络中消失
8. redis的基本数据类型
    - string\list\hash\sort\zsort
9. hash的实现原理
    - 哈希表本质是一种(key,value)结构
    - 由此我们可以联想到，能不能把哈希表的key映射成数组的索引index呢？
    - 如果这样做的话那么查询相当于直接查询索引，查询时间复杂度为O(1)
    - 其实这也正是当key为int型时的做法 将key通过某种做法映射成index，从而转换成数组结构
10. 实现步骤
    - 使用hash算法计算key值对应的hash值h(默认用key对应的hashcode进行计算(hashcode默认为key在内存中的地址)),得到hash值
    - 计算该(k,v)对应的索引值index ，索引值的计算公式为 index = (h % length) length为数组长度
    - 储存对应的(k,v)到数组中去，从而形成a[index] = node<k,v>,如果a[index]已经有了结点即可能发生碰撞，那么需要通过开放寻址法或拉链法(Java默认实现）解决冲突
10. mysql的索引用什么实现的？
    - btree
11. b+tree的实现原理
12. 复习掌握常用排序算法的复杂度及实现原理

#### 网上字节面试题
##### Redis 部分
- Redis的应用场景
- Redis支持的数据类型（必考）
- zset跳表的数据结构（必考）
- Redis的数据过期策略（必考）
- Redis的LRU过期策略的具体实现
- 如何解决Redis缓存雪崩，缓存穿透问题
- Redis的持久化机制（必考）
- Redis的管道pipeline

##### Mysql 部分
- 事务的基本要素
- 事务隔离级别
- 如何解决事务的并发问题(脏读，幻读)？
- MVCC多版本并发控制？
- binlog,redolog,undolog都是什么，起什么作用？
- InnoDB的行锁/表锁？
- myisam和innodb的区别，什么时候选择myisam？
- 为什么选择B+树作为索引结构？
- 索引B+树的叶子节点都可以存哪些东西？
- 查询在什么时候不走（预期中的）索引？
- sql如何优化?
- explain是如何解析sql的？
- order by原理
