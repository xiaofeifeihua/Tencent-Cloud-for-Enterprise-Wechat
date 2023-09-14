package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

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

	// 初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	// 设置初始的seq值为0
	var seq uint64 = 0

	// 创建文件夹
	dirPath, _ := os.Getwd()
	dirPath = path.Join(dirPath, "downloaded_images")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
	}

	for {
		// 同步消息
		chatDataList, err := client.GetChatData(seq, 100, "", "", 3)
		if err != nil {
			fmt.Printf("消息同步失败：%v \n", err)
			return
		}

		// 如果没有更多的消息，退出循环
		if len(chatDataList) == 0 {
			break
		}

		for _, chatData := range chatDataList {
			// 更新seq值以拉取更多的消息
			if chatData.Seq > seq {
				seq = chatData.Seq
			}

			// 消息解密
			chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
			if err != nil {
				fmt.Printf("消息解密失败：%v \n", err)
				return
			}

			if chatInfo.Type == "image" {
				image := chatInfo.GetImageMessage()
				sdkfileid := image.Image.SdkFileID

				isFinish := false
				buffer := bytes.Buffer{}
				index_buf := ""
				for !isFinish {
					// 获取媒体数据
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
				fileName := fmt.Sprintf("image_%d.png", time.Now().UnixNano())
				filePath := path.Join(dirPath, fileName)
				err := ioutil.WriteFile(filePath, buffer.Bytes(), 0666)
				if err != nil {
					fmt.Printf("文件存储失败：%v \n", err)
					return
				}
			}
		}
	}
}
