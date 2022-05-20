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

<hr>

### 参数说明

```json
{
  "name":"我的mac", //名字
  "uptime":1225456, //已开机时长
  "network_rx":5182, //实时入口流量
  "network_tx":8758, //实时出口流量
  "network_in":70191172528, //总入口流量
  "network_out":125968095215, //总出口流量
  "cpu":"10*Apple M1 Pro", //CPU数量型号
  "cpu_used_percent":19.806517300776,//cpu使用百分比
  "memory_total":17179869184,//ram总大小
  "memory_used_percent":81.92014694213867,//ram使用百分比
  "swap_total":5368709120,//swap总大小
  "swap_used_percent":71.9287109375,//swap使用百分比
  "hdd_total":494384795648,//硬盘总大小
  "hdd_used_percent":92.67290554263093,//硬盘使用百分比
  "rxp_now":26,//接收包实时速率
  "txp_now":24,//发送包实时速率
  "cpu_version":"21.4.0", //内核版本
  "cpu_arch":"arm64", //cpu架构
  "rxp":89273231, //接受包总量
  "txp":123814486, //发送包总量
  "tcp_num":124, //tcp连接数
  "udp_num":8 //udp连接数
}
```