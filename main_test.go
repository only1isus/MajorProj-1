package main

import (
	"os"
	"testing"

	context "golang.org/x/net/context"

	notification "github.com/azziwarlock21/MajorProj/proto"
	"google.golang.org/grpc"
)

func startClient(t *testing.T) {
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		t.Log("cannot start the server on this port")
	}
	// create new instance of context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nc := notification.NewNotifyClient(conn)
	nParams := &notification.Params{
		Msg:      "now working with grpc!",
		Reciever: os.Getenv("RECIPIENT_NUM"),
	}

	resp, err := nc.Send(ctx, nParams)
	if err != nil {
		t.Log(err)
	}
	if !resp.Success {
		t.Log("sending message failed.")
		t.Fail()
	}
}
func TestMain(t *testing.T) {
	startClient(t)
}
