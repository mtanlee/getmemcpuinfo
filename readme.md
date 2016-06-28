GetMemCpuInfo tool使用说明

```

1.配置：

在Centos 7下开发的，
go version go1.5.4 linux/amd64,
配置路径为/etc/GetInfoConf/getinfo.conf，
配置文件内容的格式为JSON格式，如下：

{
  "Ak": "",
  "Sk": "",
  "Action": "",
  "DBInstanceId": "",
  "Key": "",
  "Token": ""
}
2.给getmemcpuinfo权限（复制到/usr/local/bin）
chmod +x 0777 getmemcpuinfo
cp getmemcpuinfo /usr/local/bin

3.命令使用
./getmemcpuinfo c
    或者 
/usr/local/bin/getmemcpuinfo c
  
./getmemcpuinfo g --addr www.xxx.com --heart 60s
    或者
/usr/local/bin/getmemcpuinfo g --addr www.xxx.com --heart 60s

```
