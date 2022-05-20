# system_query
golang基于gopsutil实现的系统探针
<hr>

## 安装

### 直接安装

从 [release](https://github.com/wuwuwu223/system_query/releases) 下载编译好的版本

### 编译安装
克隆仓库到本地编译
```shell
git clone https://github.com/wuwuwu223/system_query
cd system_query
go build
```
<hr>

## 使用

```shell
./system_query -s 标识名 -i 出口网卡名
```
<hr>

### 定制化

修改源码 ```main.go``` 第63行左右 ```fmt.Println(string(str))``` 实现提交到主控服务器
