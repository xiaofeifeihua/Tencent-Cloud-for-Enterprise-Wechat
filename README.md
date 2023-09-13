## Tencent-Cloud-for-Enterprise-Wechat
在腾讯云部署的企业微信会话内容存档

### 官方文档地址
https://open.work.weixin.qq.com/api/doc/90000/90135/91774

### 环境部署：
腾讯云香港主机，登录后切换成root账号 sudo su
1.安装Golang环境,官网下载并解压缩在/usr/go/go文件夹,在/etc/profile 最后一行加上环境变量：export PATH=$PATH:/usr/go/go/bin
2.检查Golang是否安装成功：go version
3.程序包地址：https://github.com/NICEXAI/WeWorkFinanceSDK
4.上传程序到/usr/wechatdata文件下
5.修改程序中example.go中的	
corpID := 企业ID，企业微信后台登录查看
corpSecret := 企业Secret值，企业微信后台登录查看
rsaPrivateKey := RSA密钥，企业微信后台设置，由第三方网站生成
6.原example只支持获取最近发送的图片并保存在云服务中，修改example.go，增加获取文本内容并存档的代码
7.运行程序 go run example.go
