package tg_api_client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/rlarlgh/telegram-api-grpc-client/grpc/telegpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	backoffSleep = 1.0 // 백오프 초기값
)

type TgApiClient struct {
	Host   string
	Port   int
	ChRecv chan *pb.ServerPushResponse
	client pb.TelegramApiClient
}

func NewTgApiClient(host string, port int) *TgApiClient {
	return &TgApiClient{
		Host:   host,
		Port:   port,
		ChRecv: make(chan *pb.ServerPushResponse),
	}
}

func (c *TgApiClient) Start(host string, port int) {
	for {
		c.startConn(host, port)
		// 지수 백오프를 사용하여 딜레이를 줍니다.
		backoffTime := time.Duration(1.5 * float64(backoffSleep))
		log.Printf("Sleeping for %v\n", backoffTime)
		time.Sleep(backoffTime)
	}
}

// tryConnect 함수는 지정된 gRPC 서버에 연결을 시도하고, 연결 객체를 반환합니다.
// 연결이 끊어지면 지수 백오프를 사용하여 재연결을 시도합니다.
func (c *TgApiClient) tryConnect(host string, port int) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	var conn *grpc.ClientConn
	var err error

	// 지수 백오프 설정
	backoffConfig := backoff.Config{
		BaseDelay:  1.0 * time.Second,  // 최소 지연 시간
		Multiplier: 1.5,                // 지연 시간 증가 인자
		MaxDelay:   60.0 * time.Second, // 최대 지연 시간
	}

	// 연결 옵션 설정
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), // TLS 설정
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoffConfig}),
		grpc.WithBlock(), // 초기 연결 시 블록킹 모드
	}

	// 연결 시도
	conn, err = grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	// 연결 상태 모니터링
	go func() {
		for {
			state := conn.GetState()
			log.Printf("Connection state: %v\n", state)
			if state == connectivity.TransientFailure {
				// 연결 상태가 TransientFailure이면 재연결 시도
				conn.ResetConnectBackoff()
			}

			// 연결 상태가 변할 때까지 대기
			if !conn.WaitForStateChange(context.Background(), state) {
				// 연결이 종료되면 루프 종료
				break
			}
		}
	}()

	return conn, nil
}

func (c *TgApiClient) startConn(host string, port int) {
	log.Println("** Try conn StartServerPush")
	// conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithConnectParams(grpc.ConnectParams{Backoff: b}))
	conn, err := c.tryConnect(host, port)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	backoffSleep = 1.0 // 커넥션 성공 시 백오프 초기화
	c.client = pb.NewTelegramApiClient(conn)

	stream, err := c.client.StartServerPush(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("openn stream error: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Println("err", err)
			break
		}
		c.ChRecv <- resp
	}
}

func (c *TgApiClient) SendMessage(botName string, text string, chatId int64, messageId int32) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 5초 타임아웃
	defer cancel()

	r, err := c.client.SendMessage(ctx, &pb.SendMessageRequest{
		BotName:          botName,
		Text:             text,
		ChatId:           chatId,
		ReplyToMessageId: messageId,
	})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	return r.GetMessage()
}

func (c *TgApiClient) SendKeyboard(botName string, text string, chatId int64, messageId int32, buttons []string, rowNum int32) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.client.SendKeyboard(ctx, &pb.SendKeyboardRequest{
		BotName:          botName,
		Text:             text,
		ChatId:           chatId,
		ReplyToMessageId: messageId,
		Buttons:          buttons,
		RowNum:           rowNum,
	})
	if err != nil {
		log.Fatalf("could not send keyboard: %v", err)
	}
	return r.GetMessage()
}
