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

```bigquery
{
"email": "12345@qq.com"
}
```

- Response Body
```json
{
  "code": 200,
  "msg": {
    "detail": "成功",
    "email_code": "878494"
  }
}
```
| 序号 | 参数  | 类型        | 简介         |
| ---- | ----- | ----------- | ------------ |
| 1    | email | varchar(33) | 用户的邮箱   |
| 2    | msg   | string      | 发送的验证码 |

[comment]: <> "### 退出登录"

[comment]: <> "- URL: /v1/api/user/logout"

[comment]: <> "- Method: DELETE"

[comment]: <> "- Request Body"

[comment]: <> "| key   | value             |"

[comment]: <> "| ----- | ----------------- |"

[comment]: <> "| email | 2101917115@qq.com |"

[comment]: <> "- Response Body"

[comment]: <> "```json"

[comment]: <> "{"

[comment]: <> "    "code": 200,"

[comment]: <> "    "msg": {"

[comment]: <> "        "detail": "成功""

[comment]: <> "    }"

[comment]: <> "}"

[comment]: <> "```"

[comment]: <> "| 序号 | 参数  | 类型        | 简介       |"

[comment]: <> "| ---- | ----- | ----------- | ---------- |"

[comment]: <> "| 1    | email | varchar&#40;33&#41; | 用户的邮箱 |"



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

- URL:/v1/api/socket/debate
- Method: POST
- Request Body
```json
{
    "yid": 1,
    "nid": 2,
    "title":"阿巴阿巴阿巴",
    "negative_content": "我不是阿巴阿巴阿巴",
    "positive_content": "我是阿巴阿巴阿巴"
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

| 序号 | 参数             | 类型   | 简介         |
| ---- | ---------------- | ------ | ------------ |
| 1    | yid              | int    | 正方的用户id |
| 2    | nid              | int    | 反方的用户id |
| 3    | title            | string | 辩论的标题   |
| 4    | negative_content | string | 正方的发言   |
| 5    | positive_content | string | 反方的发言   |

### 查询辩论场次的数据

- URL: v1/api/socket/debate/:id
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
        "data": [
            "dGl0bGU=",
            "6Zi/5be06Zi/5be06Zi/5be0",
            "cG9zaXRpdmVfY29udGVudA==",
            "5oiR5piv6Zi/5be06Zi/5be06Zi/5be0",
            "bmVnYXRpdmVfY29udGVudA==",
            "5oiR5LiN5piv6Zi/5be06Zi/5be06Zi/5be0",
            "eWlk",
            "MQ==",
            "bmlk",
            "Mg==",
            "dGltZQ==",
            "MTYxODYxODYyNA=="
        ],
        "detail": "成功"
    }
}
```

| 序号 | 参数 | 类型 | 简介         |
| ---- | ---- | ---- | ------------ |
| 1    | id   | int  | 辩论场次的id |

返回状态码表

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

通信原理图：
![img.png](photo/img.png)
在socket通信服务端设置一个注册服务中心，一个广播中心，服务中心用于注册匹配进来的客户，通过广播中心进行一个消息的传递。