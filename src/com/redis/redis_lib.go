package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var client *redis.Client

//连接redis服务端
func Init() {
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0, //redis默认有0-15共16个数据库，这里设置操作索引为0的数据库
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	if pong != "PONG" {
		log.Fatal("客户端连接redis服务端失败")
	} else {
		fmt.Println("客户端已成功连接至redis服务端")
	}
}

//string类型数据操作
//redis命令：set key val
func Set(key, val string) {
	//有效期为0表示不设置有效期，非0表示经过该时间后键值对失效
	result, err := client.Set(key, val, 0).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：get key
func Get(key string) {
	val, err := client.Get(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：mset key1 val1 key2 val2 key3 val3 ...
func Mset(key1, val1, key2, val2, key3, val3 string) {
	//以下三种方式都可以，习惯于对象操作的我优先选择第三种
	//result,err := client.MSet(key1,val1,key2,val2,key3,val3).Result()
	//result,err := client.MSet([]string{key1,val1,key2,val2,key3,val3}).Result()
	result, err := client.MSet(map[string]interface{}{key1: val1, key2: val2, key3: val3}).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：mget key1 key2 key3 ...
func Mget(key1, key2, key3 string) {
	vals, err := client.MGet(key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：del key1 key2 key3 ...
func Del(key1, key2, key3 string) {
	result, err := client.Del(key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：getrange key start end
func Getrange(key string, start, end int64) {
	val, err := client.GetRange(key, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：strlen key
func Strlen(key string) {
	len, err := client.StrLen(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：setex key time val
func Setex(key, val string, expire int) {
	//time.Duration其实也是int64，不过是int64的别名罢了，但这里如果expire使用int64也无法与time.Second运算，
	//因为int64和Duration虽然本质一样，但表面上属于不同类型，go语言中不同类型是无法做运算的
	result, err := client.Set(key, val, time.Duration(expire)*time.Second).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//redis命令：append key val
func Append(key, val string) {
	//将val插入key对应值的末尾，并返回新串长度
	len, err := client.Append(key, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：exists key
func Exists(key string) {
	//返回1表示存在，0表示不存在
	isExists, err := client.Exists(key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isExists)
}

//hash类型数据操作
//redis命令：hset hashTable key val
func Hset(hashTable, key, val string) {
	isSetSuccessful, err := client.HSet(hashTable, key, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	//如果键存在这返回false，如果键不存在则返回true
	fmt.Println(isSetSuccessful)
}

//redis命令：hget hashTable key
func Hget(hashTable, key string) {
	val, err := client.HGet(hashTable, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：hmset hashTable key1 val1 key2 val2 key3 val3 ...
//该函数本身有问题，只插入一个键值对的话相当于hset，可以成功
//如果插入一个以上的键值对则会报错：ERR wrong number of arguments for 'hset' command
//且go-redis官方本身也不推荐是用该函数
//func hmset(hashTable,key1,val1,key2,val2,key3,val3 string){
//	_,err := client.HMSet(hashTable,key1,val1,key2,val2,key3,val3).Result()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//redis命令：hmget hashTable key1 key2 key3 ...
func Hmget(hashTable, key1, key2, key3 string) {
	vals, err := client.HMGet(hashTable, key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hdel hashTable key1 key2 key3 ...
func Hdel(hashTable, key1, key2, key3 string) {
	//返回1表示删除成功，返回0表示删除失败
	//只要至少有一个被删除则返回1（不存在的键不管），一个都没删除则返回0（不存在的则也算没删除）
	n, err := client.Del(hashTable, key1, key2, key3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：hgetall hashTable
func Hgetall(hashTable string) {
	vals, err := client.HGetAll(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hexists hashTable key
func Hexists(hashTable, key string) {
	isExists, err := client.HExists(hashTable, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isExists)
}

//redis命令：hlen hashTable
func Hlen(hashTable string) {
	len, err := client.HLen(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：hkeys hashTable
func Hkeys(hashTable string) {
	keys, err := client.HKeys(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range keys {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：hvals hashTable
func Hvals(hashTable string) {
	vals, err := client.HVals(hashTable).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//list类型数据操作
//redis命令：lpush mylist val1 val2 val3 ...
func Lpush(mylist, val1, val2, val3 string) {
	//返回列表的总长度（即有多少个元素在列表中）
	n, err := client.LPush(mylist, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：rpush mylist val1 val2 val3 ...
func Rpush(mylist, val1, val2, val3 string) {
	//返回列表的总长度（即有多少个元素在列表中）
	n, err := client.RPush(mylist, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：lpop mylist
func Lpop(mylist string) {
	//返回被删除的值
	val, err := client.LPop(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：rpop mylist
func Rpop(mylist string) {
	//返回被删除的值
	val, err := client.RPop(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：lrem mylist count val
func Lrem(mylist, val string, count int64) {
	//返回成功删除的val的数量
	n, err := client.LRem(mylist, count, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：ltrim mylist start end
func Ltrim(mylist string, start, end int64) {
	//返回状态（OK）
	status, err := client.LTrim(mylist, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

//redis命令：lset mylist index val
func Lset(mylist, val string, index int64) {
	status, err := client.LSet(mylist, index, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

//redis命令：lindex mylist index
func Lindex(mylist string, index int64) {
	//通过索引查找字符串
	val, err := client.LIndex(mylist, index).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：lrange mylist start end
func Lrange(mylist string, start, end int64) {
	vals, err := client.LRange(mylist, start, end).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：llen mylist
func Llen(mylist string) {
	len, err := client.LLen(mylist).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//无序集合set类型数据操作
//redis命令：sadd myset val1 val2 val3 ...
func Sadd(myset, val1, val2, val3 string) {
	n, err := client.SAdd(myset, val1, val2, val3).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：srem myset val
func Srem(myset, val string) {
	//删除集合中的值并返回其索引
	index, err := client.SRem(myset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(index)
}

//redis命令：spop myset
func Spop(myset string) {
	//随机删除一个值并返回
	val, err := client.SPop(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：smembers myset
func Smembers(myset string) {
	vals, err := client.SMembers(myset).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：scard myset
func Scard(myset string) {
	len, err := client.SCard(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

//redis命令：sismember myset val
func Sismember(myset, val string) {
	//判断值是否为集合中的成员
	isMember, err := client.SIsMember(myset, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isMember)
}

//redis命令：srandmember myset count
func Srandmembers(myset string, count int64) {
	vals, err := client.SRandMemberN(myset, count).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//该函数是上一个函数在只随机取一个元素的情况
func Srandmember(myset string) {
	val, err := client.SRandMember(myset).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

//redis命令：smove myset myset2 val
func Smove(myset, myset2, val string) {
	isSuccessful, err := client.SMove(myset, myset2, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isSuccessful)
}

//redis命令：sunion myset myset2 ...
func Sunion(myset, myset2 string) {
	vals, err := client.SUnion(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sunionstore desset myset myset2 ...
func Sunionstore(desset, myset, myset2 string) {
	//返回新集合的长度
	n, err := client.SUnionStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：sinter myset myset2 ...
func Sinter(myset, myset2 string) {
	vals, err := client.SInter(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sinterstore desset myset myset2 ...
func Sinterstore(desset, myset, myset2 string) {
	n, err := client.SInterStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

//redis命令：sdiff myset myset2 ...
func Sdiff(myset, myset2 string) {
	vals, err := client.SDiff(myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sdiffstore desset myset myset2 ...
func Sdiffstore(desset, myset, myset2 string) {
	n, err := client.SDiffStore(desset, myset, myset2).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
