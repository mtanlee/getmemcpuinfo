GetMemCpuInfo tool使用说明

```

1.配置：

使用的系统为Centos 7，配置路径为/etc/yunfanconf/yunfan.conf，
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
  
./getmemcpuinfo g --addr transfer.ops.yunfancdn.com --heart 60s
    或者
/usr/local/bin/getmemcpuinfo g --addr transfer.ops.yunfancdn.com --heart 60s

```