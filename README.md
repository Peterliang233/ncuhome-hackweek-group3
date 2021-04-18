# 第三组后端代码

## api文档
- [API 文档](#api文档)
- [1.用户登录](#用户登录)
- [2.用户注册](#用户注册)
- [3.邮箱验证](#邮箱验证)
- [4.修改用户信息](#修改用户信息)
- [5.修改用户密码](#修改用户密码)  
- [6.上传用户头像](#上传用户头像)
- [7.查看用户自身信息](#查看用户自身信息)  
- [8.接收辩论场次具体数据](#接收辩论场次具体数据)
- [9.查询辩论场次的数据](#查询辩论场次的数据)
- [10.查询所有辩论的数据](#查询所有辩论的数据)
- [11.获取完成的辩论场](#获取完成的辩论场)
- [12.获取将来的辩论场](#获取将来的辩论场)
- [13.选择正方](#选择正方)
- [14.选择反方](#选择反方)
- [15.添加辩题](#添加辩题)  
- [返回状态码表](#返回状态码表)
- [通信原理图](#通信原理图)

### 用户登录
- RUL: v1/api/login
- Method: POST
- Request Body

```json
{
  "email": "12345@qq.com",
  "password": "3607812001lyp",
  "remember_password": false
}
```

- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjIxMDc5MTcxMTVAcXEuY29tIiwiZXhwIjoxNjE4NDE2OTMzLCJpc3MiOiJwZXRlciJ9.rv31MRiM-oPbCSE73clMYjScZ9FM8dSeS2Jh7bYcXmw"
    }
}
```
| 序号 | 参数     | 类型        | 简介            |
| ---- | -------- | ----------- | --------------- |
| 1    | email    | varchar(33) | 用户邮箱        |
| 2    | password | varchar(33) | 用户密码        |
| 3    | token    | string      | 用户登录的token |

### 用户注册

- RUL: v1/api/registry
- Method: POST
- Request Body
```json
{
    "email": "2107917115@qq.com",
    "password":"3607812001lyp",
    "code":"341553"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功"
    }
}
```

| 序号 | 参数     | 类型        | 简介             |
| ---- | -------- | ----------- | ---------------- |
| 1    | email    | varchar(33) | 用户邮箱         |
| 2    | password | varchar(33) | 用户密码         |
| 3    | code     | varchar(10) | 用户收到的验证码 |



### 邮箱验证

- URL: v1/api/verify
- Method: POST
- Request Body
```json
{
  "username": "123456@qq.com"
}
```
- Response Body
```json
{
  "code": 200,
  "msg": {
    "detail": "成功",
    "email_code": "828208"
  }
}
```

### 修改用户信息

- URL:/v1/api/user/info
- Method: PUT
- Request Body
```json
{
    "phone": "123213132",
    "username": "peterliang"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功",
        "phone": "123213132",
        "username": "xiaoliang"
    }
}
```

| 序号 | 参数     | 类型        | 简介       |
| ---- | -------- | ----------- | ---------- |
| 1    | phone    | varchar(30) | 用户的电话 |
| 2    | username | varchar(30) | 用户名     |



### 查看用户自身信息

- URL:/v1/api/user/info
- Method: GET
- Response Body
```json
{
  "code": 200,
  "msg": {
    "data": {
      "username": "xiaoliang",
      "score": "0",
      "img": "",
      "title": "学录"
    },
    "detail": "成功"
  }
}
```
| 序号 | 参数     | 类型   | 简介         |
| ---- | -------- | ------ | ------------ |
| 1    | username | string | 用户名       |
| 2    | score    | string | 用户当前分数 |
| 3    | img      | string | 用户头像     |
| 4    | title    | string | 用户头衔     |

### 上传用户头像

- URL: v1/api/user/upload
- Method: POST
- Response Body
```json
{
    "code": 200,
    "msg": {
        "data": "",
        "detail": "upload success"
    }
}
```

| 序号 | 参数 | 类型 | 简介       |
| ---- | ---- | ---- | ---------- |
| 1    | file | file | 上传的文件 |



### 修改用户密码

- URL:v1/api/user/pwd
- Method: PUT
- Request Body
```json
{
    "email": "2107917115@qq.com",
    "old_password": "3607812001lyp",
    "new_password": "123456abc",
    "check_new_password": "123456abc"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功"
    }
}
```

| 序号 | 参数               | 类型        | 简介             |
| ---- | ------------------ | ----------- | ---------------- |
| 1    | email              | varchar(33) | 用户的邮箱       |
| 2    | old_password       | varchar(30) | 用户的旧密码     |
| 3    | new_password       | varchar(30) | 用户的新密码     |
| 4    | check_new_password | varchar(30) | 确认用户的新密码 |


### 接收辩论场次具体数据

- URL:/v1/api/socket/one
- Method: POST
- Request Body
```json
{
  "positive_username":"张三",
  "negative_username": "李四",
  "title":"阿巴阿巴阿巴",
  "negative_content": "我不是阿巴阿巴阿巴",
  "positive_content": "我是阿巴阿巴阿巴",
  "begin_time": "2021-04-17 23:59:23"
}
```
- Response Body
```json
{
  "code": 200,
  "msg": {
    "data": {
      "id": 1,
      "positive_username": "张三",
      "negative_username": "李四",
      "title": "阿巴阿巴阿巴",
      "negative_content": "我不是阿巴阿巴阿巴",
      "positive_content": "我是阿巴阿巴阿巴",
      "begin_time": "2021-04-17 23:59:23"
    },
    "detail": "成功"
  }
}
```

| 序号 | 参数             | 类型   | 简介         |
| ---- | ---------------- | ------ | ------------ |
| 1    | positive_username| string | 正方的用户id |
| 2    | negative_username| string    | 反方的用户id |
| 3    | title            | string | 辩论的标题   |
| 4    | negative_content | string | 正方的发言   |
| 5    | positive_content | string | 反方的发言   |
| 6    | begin_time       | string | 辩论开始的时间|

### 查询辩论场次的数据

- URL: v1/api/socket/record/:id
- Method: GET
- Request Body

| key  | value |
| ---- | ----- |
| id   | 1     |

- Response Body
```json
{
  "code": 200,
  "msg": {
    "data": {
      "negative_content": "我不是阿巴阿巴阿巴",
      "negative_username": "李四",
      "nid": "2",
      "positive_content": "我是阿巴阿巴阿巴",
      "positive_username": "张三",
      "time": "2021-04-17 23:59:23",
      "title": "阿巴阿巴阿巴",
      "yid": "1"
    },
    "detail": "成功"
  }
}
```

| 序号 | 参数              | 类型   | 简介           |
| ---- | ----------------- | ------ | -------------- |
| 1    | negative_content  | string | 辩论反方的记录 |
| 2    | negative_username | string | 反方的用户名   |
| 3    | nid               | string | 反方的用户id   |
| 4    | positive_content  | string | 正方的记录     |
| 5    | positive_username | string | 正方的记录     |
| 6    | yid               | string | 正方的id       |
| 7    | time              | string | 辩论的开始时间 |
| 8    | title             | string | 辩论标题       |

### 获取所有辩论记录
- URL: /v1/api/debate/records 
- Method: GET
- Request Body
```json
{
    "page_num": 3,
    "page_size":2
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "data": [
            {
                "id": 5,
                "title": "我是阿巴",
                "positive_username": "peter",
                "negative_username": "peter",
                "begin_time": "2020-01-02 12:12:12"
            },
            {
                "id": 6,
                "title": "我是阿巴",
                "positive_username": "",
                "negative_username": "peter",
                "begin_time": "2020-01-02 12:12:12"
            }
        ],
        "detail": "成功"
    }
}
```

| 序号 | 参数      | 类型 | 简介     |
| ---- | --------- | ---- | -------- |
| 1    | page_size | int  | 分页大小 |
| 2    | page_num  | int  | 分页值   |

### 选择正方

- URL: /v1/api/debate/pos
- Method: POST
- Request Body
```json
{
    "title": "我是阿巴",
    "positive_username": "peter",
    "begin_time":"2020-01-02 12:12:12"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功",
        "username": "peter"
    }
}
```
| 序号 | 参数              | 类型   | 简介             |
| ---- | ----------------- | ------ | ---------------- |
| 1    | title             | string | 论题             |
| 2    | positive_username | string | 正方用户名       |
| 3    | begin_time        | string | 辩论场次开始时间 |

### 选择反方

- URL: /v1/api/debate/neg
- Method: POST
- Request Body
```json
{
    "title": "我是阿巴",
    "negative_username": "peter",
    "begin_time":"2020-01-02 12:12:12"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功",
        "username": "peter"
    }
}
```

| 序号 | 参数              | 类型   | 简介       |
| ---- | ----------------- | ------ | ---------- |
| 1    | title             | string | 论题       |
| 2    | negative_username | string | 反方用户名 |
| 3    | begin_time        | string | 开始时间   |

### 添加辩题

- URL: /v1/api/debate/add
- Method: POST
- Request Body
```json
{
    "title": "我是阿巴",
    "begin_time": "2021-04-17 23:59:23"
}
```
- Response Body
```json
{
    "code": 200,
    "msg": {
        "data": {
            "id": 0,
            "title": "我是阿巴",
            "positive_username": "",
            "negative_username": "",
            "begin_time": "2021-04-17 23:59:23"
        },
        "detail": "成功"
    }
}
```

| 序号 | 参数       | 类型   | 简介         |
| ---- | ---------- | ------ | ------------ |
| 1    | title      | string | 论题         |
| 2    | begin_time | string | 辩论开始时间 |

### 返回状态码表

| 参数名               | 数字码 | 简介               |
| -------------------- | ------ | ------------------ |
| Success              | 200    | 成功               |
| Error                | 500    | 失败               |
| InvalidToken         | 1001   | 非法的token        |
| TokenNotExist        | 1002   | token错误          |
| TokenError           | 1003   | 请求头中的auth为空 |
| AuthEmpty            | 1004   | token不存在        |
| TokenRunTimeError    | 1005   | token过期          |
| ErrRequest           | 2001   | 请求错误           |
| ErrParameter         | 2002   | 请求参数错误       |
| ErrInfoNotFound      | 3001   | 未查找到相关信息   |
| ErrDatabaseFound     | 3002   | 数据库查找错误     |
| ErrRedisCached       | 3003   | redis存储错误      |
| ErrUserNameUsed      | 4001   | 用户名已存在       |
| ErrUserEmailUsed     | 4002   | 用户邮箱已存在     |
| ErrUserPhoneUsed     | 4003   | 用户电话已存在     |
| ErrPassword          | 4004   | 用户密码错误       |
| ErrPhoneNotExist     | 4005   | 号码不存在         |
| ErrPasswordDifferent | 4006   | 密码不一致         |
| ErrEmailNotExist     | 4007   | 邮箱不存在         |
| ErrEmailCode         | 5001   | 邮箱验证码错误     |

### 通信原理图
![img.png](photo/img.png)
在socket通信服务端设置一个注册服务中心，一个广播中心，服务中心用于注册匹配进来的客户，通过广播中心进行一个消息的传递。