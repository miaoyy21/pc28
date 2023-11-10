package server

import (
	"context"
	"fmt"
	"g28/hdo"
	pb "g28/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
)

type GoldService struct {
	pb.UnimplementedGoldServiceServer
}

type QGoldRequest struct {
	Unix      string `json:"unix"`
	KeyCode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	UserId    string `json:"userid"`
	Token     string `json:"token"`
}

type QGoldResponse struct {
	Status int `json:"status"`
	Data   struct {
		GoldEggs string `json:"goldeggs"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func (s *GoldService) Gold(ctx context.Context, r *pb.GoldRequest) (*pb.GoldResponse, error) {
	req := QGoldRequest{
		Unix:      r.GetUnix(),
		KeyCode:   r.GetKeyCode(),
		PType:     r.GetPType(),
		DeviceId:  r.GetDeviceId(),
		ChannelId: r.GetChannelId(),
		UserId:    r.GetUserId(),
		Token:     r.GetToken(),
	}

	var resp QGoldResponse

	err := hdo.Do(r.GetOrigin(), r.GetCookie(), r.GetUserAgent(), r.GetUrl(), req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	sGold := strings.ReplaceAll(resp.Data.GoldEggs, ",", "")
	iGold, err := strconv.Atoi(sGold)
	if err != nil {
		return nil, err
	}

	return &pb.GoldResponse{Gold: int64(iGold)}, nil
}

func gGold(port string) {
	server := grpc.NewServer()
	pb.RegisterGoldServiceServer(server, &GoldService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Gold net.Listen() Failure : %s \n", err.Error())
	}

	if err := server.Serve(l); err != nil {
		log.Fatalf("Gold server.Serve() Failure : %s \n", err.Error())
	}
}
