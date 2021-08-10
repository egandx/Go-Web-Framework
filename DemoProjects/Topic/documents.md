## API

GET /v1/topics  默认显示所有话题列表

GET /v1/topics?username=XXX  显示用户发布的帖子

GET /v1/topics/123  显示帖子ID为123的详细内容

POST /v1/topics  外加JSON参数，即可进行帖子的新增（注意！这是一个需要登陆权限才能使用的）

DELETE /v1/topics/123  删除帖子（注意！这是一个需要登陆权限才能使用的）

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

```html
验证器来源在于一个第三方库:
https://github.com/go-playground/validator
```

