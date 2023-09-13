package main

import (
	"fmt"
	"os"

	"github.com/NICEXAI/WeWorkFinanceSDK"
)

func main() {
	corpID := "xxx"
	corpSecret := "xxx"
	rsaPrivateKey := `
-----BEGIN RSA PRIVATE KEY-----
xxx
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

	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}

		if chatInfo.Type == "text" {
			textMessage := chatInfo.GetTextMessage()
			message := textMessage.Text.Content // Corrected this line
			sender := textMessage.From

			// Construct the string to save in the .txt file
			saveString := fmt.Sprintf("Sender: %s\nMessage: %s\n\n", sender, message)

			// Append the message to a .txt file
			f, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error opening the .txt file: %v\n", err)
				return
			}
			defer f.Close()
			if _, err := f.WriteString(saveString); err != nil {
				fmt.Printf("Error writing to the .txt file: %v\n", err)
				return
			}
		}
	}
}
