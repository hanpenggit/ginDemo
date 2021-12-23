# ginDemo

### 添加了自定义jwt的权限校验
### 端口和日志的位置，在./config.yaml配置文件中

### 代码结构：
```
ginDemo
├── app
│   ├── role
│   │   ├── handler.go
│   │   └── router.go
│   └── user
│       ├── handler.go
│       └── router.go
├── config.yaml
├── go.mod
├── go.sum
├── main.go
├── model
│   └── model.go
├── README.md
├── routers
│   └── routers.go
├── service
│   └── publicService.go
└── utils
    ├── jwtUtil.go
    └── logUtil.go
```


### 登录接口
#### http://localhost:8085/login    
#### Content-Type : application/json
```
{
    "username":"admin",
    "password":"admin"
}

返回值：

{
    "code": 200,
    "message": "登录成功",
    "data": {
        "exp": 3600,
        "token": "xxxxx"
    }
}
```

### 获取用户详情
#### http://localhost:8085/user/detail
#### 添加 header
```
Authorization=Bearer xxxxx

返回值：

{
    "code": 200,
    "message": "获取用户信息成功",
    "data": {
        "exp": 3586,
        "permission": [
            "/role/a",
            "/user/a"
        ],
        "role": [
            "admin",
            "user"
        ],
        "username": "admin"
    }
}
```
