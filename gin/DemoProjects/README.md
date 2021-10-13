# 项目结构
```
.
├── Data
│   ├── MySQL
│   │   └── gin.sql
│   └── img
│       └── tree.png
├── README.md
├── Test
└── Topic
    ├── main.go
    └── src
        ├── Dao.go
        ├── Decorator.go
        ├── MyInit.go
        ├── MyRedis.go
        ├── MyValidator.go
        ├── TopicDao.go
        └── TopicModel.go
```
/Data存放数据，例如数据库表和数据，方便更换设备时，数据库的迁移。

/Test存放测试类，暂时还没写测试，后面再说吧

/Topic存放项目的主要代码



# API

GET /v1/topics  默认显示所有帖子的详细内容

GET /v1/topics/search?username=A  查找用户A发布的帖子

GET /v1/topics/123  显示帖子ID为123的详细内容

POST /v1/topics  外加JSON参数，即可进行帖子的新增（注意！这是一个需要登陆权限才能使用的）

POST /v1/mtopics  和上面对比，一次性提交多条帖子

DELETE /v1/topics/123  删除帖子（注意！这是一个需要登陆权限才能使用的）

# API交互示例
#### [send]
```http://localhost:8080/v1/topics```

#### [respond]
```json
[
  {
    "id": 1,
    "title": "兴趣电商助力经济发展",
    "stitle": "兴趣电商",
    "ip": "192.168.251.11",
    "score": 12,
    "url": "https://192.168.251.11/1",
    "date": "2021-08-12T15:32:41+08:00",
    "username": "jerry"
  },
  {
    "id": 2,
    "title": "阿里蚂蚁捐款7000万驰援山西",
    "stitle": "阿里捐款",
    "ip": "192.168.251.12",
    "score": 34,
    "url": "https://192.168.251.11/2",
    "date": "2021-08-12T15:32:55+08:00",
    "username": "tom"
  },
  {
    "id": 3,
    "title": "美团被罚，反垄断监管规则更加清晰",
    "stitle": "美团被罚",
    "ip": "192.168.251.13",
    "score": 56,
    "url": "https://192.168.251.11/3",
    "date": "2021-08-12T15:33:14+08:00",
    "username": "玛卡巴卡"
  }
]
```
#### [send]
```http://localhost:8080/v1/topics/2```

#### [respond]
```json
{
    "id": 2,
    "title": "阿里蚂蚁捐款7000万驰援山西",
    "stitle": "阿里捐款",
    "ip": "192.168.251.12",
    "score": 34,
    "url": "https://192.168.251.11/2",
    "date": "2021-08-12T15:32:55+08:00",
    "username": "tom"
}
```

#### [send]
```http://localhost:8080/v1/topics/search?username=玛卡巴卡```

#### [respond]
```json
{
    "Page_size": 10,
    "Size": 1,
    "Topic_list": [
        {
            "id": 3,
            "title": "美团被罚，反垄断监管规则更加清晰",
            "stitle": "美团被罚",
            "ip": "192.168.251.13",
            "score": 56,
            "url": "https://192.168.251.11/3",
            "date": "2021-08-12T15:33:14+08:00",
            "username": "玛卡巴卡"
        }
    ],
    "Total_page": 1
}
```

#### [send]
```http://localhost:8080/v1/topics?token=124```
```
auth user/password=admin/123
```
```json
{"title":"这是一条测试插入的记录","ip":"192.168.251.123","score":1122,"username":"admin"}
```

#### [respond]
```json
"OK"
```


#### [send]
```http://localhost:8080/v2/mtopics?token=ssss```
```
auth user/password=admin/123
```
```json
{
    "topics":    [
{
    "title": "联想控股辟谣柳传志1亿年薪：去年已不领薪",
    "stitle": "联想",
    "ip": "192.168.251.11",
    "score": 123,
    "url": "https://192.168.251.11/57",
    "username": "admin"
},
{
    "title": "《长津湖》超过《复仇者联盟4》，成中国影史票房第6",
    "stitle": "长津湖",
    "ip": "192.168.251.11",
    "score": 965,
    "url": "https://192.168.251.11/123",
    "username": "admin"
}
    ],
    "size":2
}
```

#### [respond]
```json
"batch insert OK"
```

#### [send]
```http://localhost:8080/v1/topics/6?token=qwqw```
```
auth user/password=admin/123
```

#### [respond]
```json
"This topic does not exist" 
or
"del ok"
```


# 项目笔记
1、标题长度必须是4--20 ```binding:"min=4,max=20"```

2、副标题和主标题不能一样 ```binding:"nefield=TopicTitle"```

3、userip必须是IPv4格式 ```binding:"ipv4"```

4、score要么不填，如果填则必须大于5分 ```binding:"omitempty,gt=5"```

```html
验证器来源在于一个第三方库:
https://github.com/go-playground/validator
```

## 参数绑定
参数绑定Model

### 1、query参数绑定

```go
type TopicQuery struct {
    Username string `form:"username" json:"username" binding:"required"`
    Page int `json:"page"`
    Pagesize int `json:"pagesize"`
}
```

```text
binding:"required" //加了这个，表示该参数必不可少，少了就报错
参数错误：
Key: 'TopicQuery.Username' Error:Field validation for 'Username' failed on the 'required' tag
```
form 决定了绑定query参数的key到底是啥

另外两个没写form 不会进行绑定

### 2、JSON参数绑定

```go
type Topic struct {
	TopicID int `json:"id" binding:"required"`
	TopicTitle string `json:"title" binding:"required"`
}
```


## 自定义验证器结合正则验证JSON参数
### 假设需求

1、请求topic详细信息时：

/topics/123（ID形式）

2、也可以是：

/topics/wodetiezi(拼音样式的URL，或为了SEO等原因)


```go
struct扩展:
	
type Topic struct {
	TopicID int `json:"id" binding:"required"`
	TopicTitle string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"nefield=TopicTitle"`
	UserIP string `json:"ip" binding:"ipv4"`
	TopicScore int `json:"score" binding:"omitempty,gt=5"`
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
}
```

## 加入组
```go
type Topics struct { //多个实体
    TopicList []Topic `json:"topics"`
    TopicListSize int `json:"size"`
}
```
#### 要求：
1、TopicList必须大于0，否则无意义。且要求小于某数字，否则服务器资源不够。

2、TopicList的长度必须和TopicListSize相等，算是一个辅助验证手段。

# 数据库和ORM

## 开启Gorm之旅
```
基础配置：
驱动：https://github.com/go-sql-driver/mysql
gorm地址：https://github.com/go-gorm/gorm
安装：
# go get -u gorm.io/gorm
# go get -u github.com/go-sql-driver/mysql
```
### 1、打印日志

v1版本的gorm，使用```db.LogMode(true)```来显示数据库的Log。

in v1:
```
func (s *DB) LogMode(enable bool) *DB {
	if enable {
		s.logMode = detailedLogMode
	} else {
		s.logMode = noLogMode
	}
	return s
}
```

And in global can do set:
```
db.LogMode(m.Debug) // m.Debug = true or false
```

v2版本的gorm,in v2 How to achieve the same effect？Silent or Error or Warn or Info? can achieve the same effect?
```
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
})
```

### 2、数据库建表

(1)第一张表topic

(2)第二张表topic_classes

翻译成模型
```cgo
type TopicClass struct { 
	ClassId int
	ClassName string
	ClassRemark string
}
```
#### 表名规则
1、根据Struct名称，默认将其改写为小写并修改为复数形式。例如：
```text
struct 名称为
1）Test，对应的数据库表名为 tests，
2）TopicClass 对应的数据库表名为 topic_classes(是一个复数)
```
gorm 也有对应的函数，使其不加复数。

in v1：
```go
db.SingularTable(true)
```

in v2：
```go
dsn := "root:12345678@/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
```

#### 指定表名或者列名
在维护历史项目的时候，很多时候没有办法去修改数据库的表名或者列名。需要我们
在程序中指定表名或者列名。

```go
type TopicClass struct {
	ClassId int
	ClassName string
	ClassRemark string
	ClassType string `gorm:"column:classtype"`
}
```

### 连接池设置

https://gorm.io/docs/generic_interface.html
```text
// Get generic database object sql.DB to use its functions
sqlDB, err := db.DB()

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns sets the maximum number of open connections to the database.
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
sqlDB.SetConnMaxLifetime(time.Hour)
```

### 数据库连接出错时，关闭Web服务的两种方式
1、使用```log.Fatal("DB初始化失败：",err)```
```go
func Fatal(v ...interface{}) {
std.Output(2, fmt.Sprint(v...))
os.Exit(1)
}
```
因为Fatal的定义中是含有退出操作的```os.Exit(1)```

2、使用信号```os.Signal```
```go
srv := &http.Server{
Addr:    ":8080",
Handler: router,
}

go func() {
// 服务连接
if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
log.Fatalf("listen: %s\n", err)
}
}()

// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
quit := make(chan os.Signal)
signal.Notify(quit, os.Interrupt)
<-quit
log.Println("Shutdown Server ...")

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := srv.Shutdown(ctx); err != nil {
log.Fatal("Server Shutdown:", err)
}
log.Println("Server exiting")
```

## Gin+Redis 简单了解

### Redis

```
Redis的下载、安装文档的地址
https://redis.io/download

Redis中文教程
https://www.redis.com.cn/tutorial.html

redigo的GitHub地址：
https://github.com/gomodule/redigo
```

### 结合Gin实现基本的Redis缓存、缓存穿透简单处理
eg API:

GET /topic/4   获取帖子ID为4的帖子

最简单的缓存方案：

1、根据ID查看Redis缓存是否有值，如果有，则取出redis的内容并返回。

2、如果没有,则从MySQL数据库中取出。取出之后，放入Redis缓存，并设置过期时间。

```go
conn := RedisDefaultPool.Get()
	defer conn.Close()
	redisKey := "topic_" + tid

	ret, err := redis.Bytes(conn.Do("get", redisKey))

	if err != nil { //redis hasn't v,goto DB
		db.Find(&topics, tid)
		retData, _ := json.Marshal(topics)

		if topics.TopicID == 0{  //DB 没有匹配到
			conn.Do("setex",redisKey,20,retData)
		}else{ //DB正常数据，缓存50s
			conn.Do("setex",redisKey,50,retData)
		}

		c.JSON(http.StatusOK,topics)
		log.Println("从数据库读取")

	} else { //缓存里有值
		json.Unmarshal(ret,&topics)
		c.JSON(http.StatusOK,topics)
		log.Println("从Redis读取")
	}
```

### 使用装饰器函数 实现redis封装
高阶函数
```go
func CacheDecortor(h gin.HandlerFunc) gin.HandlersChain{
    return func(context *gin.Context) {
    }
}

// router
v1.GET("/:topic_id", CacheDecortor(GetTopicDetail))
```

```go
// Add parameters to make it more generic
func CacheDecorator(h gin.HandlerFunc,param string,redKeyPattern string,empty interface{}) gin.HandlerFunc {
    return func(context *gin.Context) {
    // redis determine
    }
}

param 是获取的参数ID,装饰器并不清楚获取的参数ID是多少
redKeyPattern 是redis中key的格式,装饰器并不清楚redis存的key是什么格式
empty 传入一个空对象,用于转化
```

# end

