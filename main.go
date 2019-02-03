package main

import (
	"fmt"
	clientLib "github.com/hackerNews/client"
	"github.com/hackerNews/data"
)

func main() {
	mobileClient := MobileClient{}
	clientResponse := clientLib.CreateClientResponse(mobileClient, mobileClient)
	go data.GetTopIds(clientResponse)
	clientLib.WaitOnBothChannels(clientResponse)
}

type MobileClient struct {
}

func (client MobileClient) Response(response string) {
	fmt.Print("RESPONSE IN MOBILE CLIENT : ", response)
}

func (client MobileClient) Error(errorResponse string) {
	fmt.Print("ERROR IN MOBILE CLIENT : ", errorResponse)
}
