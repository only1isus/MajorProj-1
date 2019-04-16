package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	//"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Set account keys & information

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Create possible message bodies
	alerts := []string{"We're good\n",
		"Temperature is under its nominal value",
		"Temperature is over its nominal value",
		"case4",
		"case5"}

	// Pack up the data for our message
	//conf := config.New()
	msgData := url.Values{}
	//fmt.Println(conf.reciever)
	reciever := os.Getenv("RECIPIENT_NUM")
	sender := os.Getenv("TWIL_NUM")
	msgData.Set("To", reciever)
	msgData.Set("From", sender)

	temp := 25
	for temp >= 22 && temp <= 29 {
		msgData.Set("Body", alerts[0])
		break
	}
	if temp < 22 {
		msgData.Set("Body", alerts[1])

	} else if temp > 35 {
		msgData.Set("Body", alerts[2])
	}

	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
