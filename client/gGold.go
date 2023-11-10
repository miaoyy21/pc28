package client

import (
	"context"
	pb "g28/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func gGold(target, cookie, userAgent, unix, keyCode, deviceId, userId, token string) (int64, error) {
	// Create a client connection to the given target with a credentials which disables transport security
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := pb.NewGoldServiceClient(conn)

	req := &pb.GoldRequest{
		Url:       conf.GoldURL,
		Origin:    conf.Origin,
		Cookie:    cookie,
		UserAgent: userAgent,
		Unix:      unix,
		KeyCode:   keyCode,
		PType:     conf.PType,
		DeviceId:  deviceId,
		ChannelId: conf.ChannelId,
		UserId:    userId,
		Token:     token,
	}

	resp, err := client.Gold(context.Background(), req)
	if err != nil {
		return 0, err
	}

	return resp.GetGold(), nil
}
