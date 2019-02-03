package main

import (
	clientLib "github.com/hackerNews/client"
	"github.com/hackerNews/data"
)

func main() {
	client := clientLib.Client{
		ErrStringChannel:  make(chan string),
		RespStringChannel: make(chan string),
	}
	client.Success = client
	client.Failure = client

	go data.GetTopIds(client)
	client.WaitOnBothChannels()
}