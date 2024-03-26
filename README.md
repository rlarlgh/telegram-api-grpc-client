# telegram-api-grpc-client

## 사용법
```
package main

import (
	"fmt"

	tagc "github.com/rlarlgh/telegram-api-grpc-client"
)

func StartGrpcClient(host string, port int) {

	client := tagc.NewTgApiClient(host, port)

	go func() {
		for recv := range client.ChRecv {
			fmt.Println("recv:", recv)
			client.SendKeyboard(recv.BotName, recv.ChatId, "버튼을 눌러주세요.", []string{"키보드1", "키보드2", "키보드3"}, 2)
			client.SendReplyKeyboard(recv.BotName, recv.ChatId, recv.MessageId, "버튼을 눌러주세요.", []string{"키보드1", "키보드2", "키보드3"}, 2)
			client.SendText(recv.BotName, recv.ChatId, "안녕하세요. 반가워요.")
			client.SendReplyText(recv.BotName, recv.ChatId, recv.MessageId, "안녕하세요. 반가워요.")
		}
	}()

	client.Start(host, port)
}

func main() {
	StartGrpcClient("localhost", 50051)
}

```