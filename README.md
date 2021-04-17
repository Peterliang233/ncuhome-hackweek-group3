# 第三组后端代码

## api文档
- [API 文档](#api文档)
- [1.用户登录](#用户登录)
- [2.用户注册](#用户注册)
- [3.邮箱验证](#邮箱验证)  
- [4.退出登录](#退出登录)  
- [5.修改用户信息](#修改用户信息)
- [6.修改用户密码](#修改用户密码)  
- [7.上传用户头像](#上传用户头像)
- [8.查看用户自身信息](#查看用户自身信息)  
- [9.接收辩论场次具体数据](#接收辩论场次具体数据)
- [10.查询辩论场次的数据](#查询辩论场次的数据)


### 用户登录
- RUL: v1/api/login
- Method: POST
- Request Body

```json
{
    "email": "2107917115@qq.com",
    "password": "3607812001lyp"
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
| key   | value         |
| ----- | ------------- |
| email | 123456@qq.com |

- Response Body
```json
{
    "msg": "651169"
}
```
| 序号 | 参数  | 类型        | 简介         |
| ---- | ----- | ----------- | ------------ |
| 1    | email | varchar(33) | 用户的邮箱   |
| 2    | msg   | string      | 发送的验证码 |

### 退出登录

- URL: /v1/api/user/logout
- Method: DELETE
- Request Body
| key   | value             |
| ----- | ----------------- |
| email | 2101917115@qq.com |

- Response Body
```json
{
    "code": 200,
    "msg": {
        "detail": "成功"
    }
}
```
| 序号 | 参数  | 类型        | 简介       |
| ---- | ----- | ----------- | ---------- |
| 1    | email | varchar(33) | 用户的邮箱 |



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