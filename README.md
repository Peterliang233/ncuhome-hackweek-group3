# 第三组后端代码

## api文档
- [API 文档](#api文档)
    + [basic Response Body]
- [1.用户登录](#用户注册)
- [2.用户注册](#用户登录)
- [3.邮箱验证](#邮箱验证)  
- [4.退出登录](#退出登录)  
- [5.修改用户信息](#修改用户信息)
- [6.修改用户密码](#修改用户密码)  
- [7.上传用户头像](#上传用户头像)


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

### 邮箱验证
- URL: v1/api/verify
- Method: POST
- Request Body
```json
"email": "123456@qq.com"
```

- Response Body
```json
{
    "msg": "651169"
}
```
### 退出登录
- URL: /v1/api/user/
- Method: DELETE
- Request Body
```json

```
- Response Body
```json

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

### 获取用户信息
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