package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"pc28/hdo"
	pb "pc28/proto"
)

type BettingService struct {
	pb.UnimplementedBettingServiceServer
}
type BettingResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type BettingRequest struct {
	Issue     string `json:"issue"`
	GoldEggs  int32  `json:"totalgoldeggs"`
	Numbers   int32  `json:"numbers"`
	Unix      string `json:"unix"`
	Keycode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	Userid    string `json:"userid"`
	Token     string `json:"token"`
}

func (s *BettingService) Betting(ctx context.Context, r *pb.BettingRequest) (*pb.BettingResponse, error) {
	req := BettingRequest{
		Issue:     r.GetIssue(),
		Unix:      r.GetUnix(),
		Keycode:   r.GetKeyCode(),
		PType:     r.GetPType(),
		DeviceId:  r.GetDeviceId(),
		ChannelId: r.GetChannelId(),
		Userid:    r.GetUserId(),
		Token:     r.GetToken(),
	}

	for i, g := range r.GetBets() {
		var resp BettingResponse

		req.Numbers = i
		req.GoldEggs = g

		err := hdo.Do(r.GetOrigin(), r.GetCookie(), r.GetUserAgent(), r.GetUrl(), req, &resp)
		if err != nil {
			return nil, err
		}

		if resp.Status != 0 {
			return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
		}
	}

	return &pb.BettingResponse{}, nil
}

func gBetting(port string) {
	server := grpc.NewServer()
	pb.RegisterBettingServiceServer(server, &BettingService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Gold net.Listen() Failure : %s \n", err.Error())
	}

	if err := server.Serve(l); err != nil {
		log.Fatalf("Gold server.Serve() Failure : %s \n", err.Error())
	}
}
