package main

import (
	"fmt"
	clientLib "github.com/hackerNews/client"
	"github.com/hackerNews/data"
)

func main() {
	mobileClient := ExternalClient{}

	internalClient := clientLib.CreateClient(mobileClient, mobileClient)
	go data.GetTopIds(internalClient)
	clientLib.WaitOnBothChannels(internalClient)
}

type ExternalClient struct {
}

func (client ExternalClient) Response(response string) {
	fmt.Print("EXTERNAL CLIENT RESPONSE : ", response)
}

func (client ExternalClient) Error(errorResponse string) {
	fmt.Print("EXTERNAL CLIENT ERROR : ", errorResponse)
}
