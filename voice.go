package main

import (
	"bytes"
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

	// 创建 voice_files 文件夹，不管消息类型
	os.MkdirAll("voice_files", os.ModePerm)
	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}

		// 检查消息类型是否为语音
		if chatInfo.Type == "voice" {
			voice := chatInfo.GetVoiceMessage()
			sdkfileid := voice.Voice.SdkFileID

			// 使用 sdkfileid 从企业微信服务器下载语音文件
			mediaData, err := client.GetMediaData("", sdkfileid, "", "", 10) // Use default values for other parameters for simplicity
			if err != nil {
				fmt.Printf("下载语音文件失败：%v \n", err)
				return
			}

			// 将下载的语音文件保存到当前目录下名为 voice_files 的文件夹中
			savePath := path.Join(".", "voice_files", sdkfileid+".mp3")
			os.MkdirAll(path.Dir(savePath), os.ModePerm)
			ioutil.WriteFile(savePath, mediaData.Data, os.ModePerm)
		}
		if chatInfo.Type == "image" {
			image := chatInfo.GetImageMessage()
			sdkfileid := image.Image.SdkFileID

			isFinish := false
			buffer := bytes.Buffer{}
			index_buf := ""
			for !isFinish {
				//获取媒体数据
				mediaData, err := client.GetMediaData(index_buf, sdkfileid, "", "", 5)
				if err != nil {
					fmt.Printf("媒体数据拉取失败：%v \n", err)
					return
				}
				buffer.Write(mediaData.Data)
				if mediaData.IsFinish {
					isFinish = mediaData.IsFinish
				}
				index_buf = mediaData.OutIndexBuf
			}
			filePath, _ := os.Getwd()
			filePath = path.Join(filePath, "test.png")
			err := ioutil.WriteFile(filePath, buffer.Bytes(), 0666)
			if err != nil {
				fmt.Printf("文件存储失败：%v \n", err)
				return
			}
			break
		}
	}
}
