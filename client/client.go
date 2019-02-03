package client

import "fmt"

type Client struct {
	RespStringChannel chan string
	ErrStringChannel  chan string
	Success           ByteResponder
	Failure           ErrorResponder
}

type ByteResponder interface {
	Response(response string)
	//Bytes(response []byte)
}

type ErrorResponder interface {
	Error(errResponse string)
	//Error(response []byte)
}

func SendErr(errorChannel chan string, dataChannel chan string, e error) {
	errorChannel <- e.Error()
	dataChannel <- ""
}

func SendData(dataChannel chan string, errorChannel chan string, data string) {
	dataChannel <- data
	errorChannel <- ""
}

// WaitOnBothChannels will wait for the bytes and error and once received it will call corresponding bytes & error responders
func (client Client) WaitOnBothChannels() {
	for i := 0; i < 2; i++ {
		select {
		case responseString := <-client.RespStringChannel:
			if responseString != "" {
				client.Success.Response(responseString)
			}
		case err := <-client.ErrStringChannel:
			// TODO handleErrorsBeforeSendingToClient(c, err)
			client.Failure.Error(err)
		}
	}
}

func (client Client) Response(response string) {
	fmt.Print("EXTERNAL CLIENT RESPONSE : ", response)
}

func (client Client) Error(errorResponse string) {
	fmt.Print("EXTERNAL CLIENT ERROR : ", errorResponse)
}
