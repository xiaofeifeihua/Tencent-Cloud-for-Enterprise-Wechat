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

6.原example只支持获取最近发送的图片并保存在云服务中，参考官方文档 https://open.work.weixin.qq.com/api/doc/90000/90135/91774 。并修改example.go，增加获取文本内容并存档的代码

7.运行程序 go run example.go

坑：腾讯云退出再登陆后会自动修改某些文件的权限，使得go无法使用，可以重新下载安装一遍

chmod -r 777 folder   给文件夹及下面所有子文件添加root权限

----------------------------------------------------------------------

修改了例子代码 - > modified_example.go , 增加保存文本消息为txt文档功能，音频保存及转录为文字/定期拉取正在开发中
