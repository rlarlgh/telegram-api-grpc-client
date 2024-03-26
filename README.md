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
			// client.SendKeyboard(recv.BotName, "키보드를 눌러주세요", recv.ChatId, recv.MessageId, []string{"키보드11", "키보드22", "키보드33"}, 2)
			// client.SendMessage(recv.BotName, "메세지를 보냅니다.", recv.ChatId, recv.MessageId)
		}
	}()

	client.Start(host, port)
}

func main() {
	StartGrpcClient("localhost", 50051)
}
```