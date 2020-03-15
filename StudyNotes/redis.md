## Redis Note

### 数据类型
#### string 
- string 类型是二进制安全的，这意味着redis的string可以包含任何数据，比如jpg图像或者序列化的对象
- redis最基本的数据类型，最大能存储512MB
#### hash
- hash是一个(key-value)对的集合
- hash是string类型的field和value的映射表，hash特别适合用于存储对象
- 用于 ： 存储、读取、修改存储对象属性
- `HMSET [key] [filed1] [value1] [field2] [value2]`
- `HGET [key] [field1]`
#### list
- list就是简单的string列表，按照插入顺序排序
- 可以添加一个元素带列表的头部或者尾部
- 用于： 消息队列、最新消息排行等(例如朋友圈的时间线)
```bash
#插入数据
lpush [key] [value1]
lpush [key] [value2]
lpush [key] [value3]
#读取数据
lrange [key] [leftIndex] [rightIndex]
```
#### set
- set是string的无序集合,不允许有重复成员
- 通过哈希表实现的集合，添加、删除、查找的复杂度都是O(1)
- 用于：
    1. 共同好友 
    2. 利用唯一性,统计访问网站的所有独立ip 
    3. 好友推荐时,根据tag求交集,大于某个阈值就可以推荐
```bash
#插入数据
sadd [key] [value1]
sadd [key] [value2]
sadd [key] [value3]
#读取数据
smembers [key]
```
#### zset
- string类型元素的集合，不允许有重复成员
- 每个元素都会关联一个double类型的分数，redis通过分数来为集合中的成员进行从小到大的排序
- zset中成员是唯一的，但是分数不可以重复
- 用法：
    1. 排行榜 
    2. 带权重的消息队列
```bash
#添加数据
zadd [key] [score1] [member1]
zadd [key] [score2] [member2]
zadd [key] [score3] [member3]
#读取数据
zrange [key] [leftIndex] [rightIndexOK]
```
