# api说明文档

包括用户验证与备忘录操作两部分

## 用户验证


### 用户注册
* 端口:localhost:8080/user/register

### 用户登录
* 端口:localhost:8080/user/login

### 用户信息获取
* 端口:localhost:8080/user/auth/:id

## 备忘录操作

### 备忘录列表获取
* 端口:localhost:8080/user/auth/:id/task

### 待办备忘录获取
* 端口:localhost:8080/user/auth/:id/task/todo

### 增
* 端口:localhost:8080/user/auth/:id/task/create
### 删
* 端口:localhost:8080/user/auth/:id/task/delete
### 改
* 端口:localhost:8080/user/auth/:id/task/update
### 查
* 端口:localhost:8080/user/auth/:id/task/search