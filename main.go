package main

import (
	"github.com/hackerNews/client"
	"github.com/hackerNews/data"
)

func main() {
	client := client.Client{
		ErrStringChannel:  make(chan string),
		RespStringChannel: make(chan string),
	}
	client.Success = client
	client.Failure = client

	go data.GetTopIds(client)
	client.WaitOnBothChannels()
}