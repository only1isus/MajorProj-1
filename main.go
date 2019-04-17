package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	notification "github.com/azziwarlock21/MajorProj/proto"
	"github.com/joho/godotenv"
)

const (
	grpcAddress = ":8083"
)

type notificationServer struct {
}

// rpc send method takes the message and the number to send and calls the API to send the message
func (nc *notificationServer) Send(ctx context.Context, params *notification.Params) (*notification.Response, error) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	sender := os.Getenv("TWIL_NUM")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", fmt.Sprintf("+%s", params.Reciever))
	msgData.Set("From", fmt.Sprintf("+%s", sender))
	msgData.Set("Body", params.Msg)

	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return &notification.Response{Success: true}, nil
	}
	return &notification.Response{Success: false}, nil
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		os.Exit(2)
	}
}

func main() {
	// Creating grpc server
	grpcServer := grpc.NewServer()
	notification.RegisterNotifyServer(grpcServer, &notificationServer{})
	conn, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to create gRPC listener: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("gRPC server started on port %s", grpcAddress)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to create gRPC serve: %v\n", err)
		os.Exit(2)
	}
}
