package data

import (
	clientLib "github.com/hackerNews/client"
	"github.com/hackerNews/network"
)

const rootUrl = "https://hacker-news.firebaseio.com/v0"

func GetTopIds(client clientLib.Client) {
	channel := make(chan network.NetworkResponse)
	path := rootUrl + "/topstories.json?"
	go network.HTTPClient.GETRequest(path, channel)
	result := <-channel

	if result.Err != nil {
		clientLib.SendErr(client.ErrStringChannel, client.RespStringChannel, result.Err)
		//SendErr(client.ErrStringChannel, result.Err)
		return
	}
	clientLib.SendData(client.RespStringChannel, client.ErrStringChannel, result.Resp)
}
