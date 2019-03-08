package test

import (
	pb "fabric-sdk-go/protos"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	ServerAddress string = "localhost:8080"
)

func CreateChannel(channelId string) (pb.StatusCode, error) {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	c := pb.NewChannelClient(conn)
	context := context.Background()
	body := &pb.CreateChannelRequest{ChannelId: channelId}

	r, err := c.CreateChannel(context, body)
	fmt.Printf("StatusCode: %s, transaction id: %s, err: %v", r.Status, r.TransactionId, err)
	return r.Status, err
}

func JoinChannel(channelId string) (pb.StatusCode, error) {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return pb.StatusCode_FAILED, err
	}

	c := pb.NewChannelClient(conn)
	context := context.Background()
	body := &pb.JoinChannelRequest{ChannelId: channelId}

	r, err := c.JoinChannel(context, body)
	fmt.Printf("StatusCode: %s, err: %v", r.Status, err)
	return r.Status, err
}