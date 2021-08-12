## API

GET /v1/topics  默认显示所有话题列表

GET /v1/topics?username=XXX  显示用户发布的帖子

GET /v1/topics/123  显示帖子ID为123的详细内容

POST /v1/topics  外加JSON参数，即可进行帖子的新增（注意！这是一个需要登陆权限才能使用的）

POST /v1/mtopics  和上面对比，一次性提交多条帖子

DELETE /v1/topics/123  删除帖子（注意！这是一个需要登陆权限才能使用的）


## 加个需求
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

    binding:"required" //加了这个，表示该参数必不可少，少了就报错
    参数错误：
    Key: 'TopicQuery.Username' Error:Field validation for 'Username' failed on the 'required' tag
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