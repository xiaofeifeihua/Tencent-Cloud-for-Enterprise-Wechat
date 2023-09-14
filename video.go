package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/NICEXAI/WeWorkFinanceSDK"
)

func main() {
	corpID := "XXX"
	corpSecret := "XXX"
	rsaPrivateKey := `
-----BEGIN RSA PRIVATE KEY-----
XXX
-----END RSA PRIVATE KEY-----
`

	//初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	//同步消息
	chatDataList, err := client.GetChatData(0, 100, "", "", 3)
	if err != nil {
		fmt.Printf("消息同步失败：%v \n", err)
		return
	}

	// 创建 video_files 文件夹，不管消息类型
	os.MkdirAll("video_files", os.ModePerm)
	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}

		// 检查消息类型是否为语音
		if chatInfo.Type == "video" {
			video := chatInfo.GetVideoMessage()
			sdkfileid := video.Video.SdkFileID

			// 使用 sdkfileid 从企业微信服务器下载语音文件
			mediaData, err := client.GetMediaData("", sdkfileid, "", "", 10) // Use default values for other parameters for simplicity
			if err != nil {
				fmt.Printf("下载视频文件失败：%v \n", err)
				return
			}

			// 将下载的语音文件保存到当前目录下名为 video_files 的文件夹中
			savePath := path.Join(".", "video_files", sdkfileid+".mp4")
			os.MkdirAll(path.Dir(savePath), os.ModePerm)
			ioutil.WriteFile(savePath, mediaData.Data, os.ModePerm)
		}

	}
}
