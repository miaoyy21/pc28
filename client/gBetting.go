package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "pc28/proto"
)

func gBetting(target, issue string, bets map[int32]int32, cookie, userAgent, unix, keyCode, deviceId, userId, token string) error {
	// Create a client connection to the given target with a credentials which disables transport security
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewBettingServiceClient(conn)

	req := &pb.BettingRequest{
		Url:       conf.BettingURL,
		Origin:    conf.Origin,
		Cookie:    cookie,
		UserAgent: userAgent,

		Issue: issue,
		Bets:  bets,

		Unix:      unix,
		KeyCode:   keyCode,
		PType:     conf.PType,
		DeviceId:  deviceId,
		ChannelId: conf.ChannelId,
		UserId:    userId,
		Token:     token,
	}

	if _, err := client.Betting(context.Background(), req); err != nil {
		return err
	}

	return nil
}
