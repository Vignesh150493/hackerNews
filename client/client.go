package client

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

// CreateClient will create client response with client implementation of bytes & error responders
func CreateClient(success ByteResponder, failure ErrorResponder) Client {
	return Client{RespStringChannel: make(chan string), ErrStringChannel: make(chan string), Success: success, Failure: failure}
}

func SendErr(errorChannel chan string, dataChannel chan string, e error) {
	errorChannel <- e.Error()
	dataChannel <- ""
}

func SendData(dataChannel chan string, errorChannel chan string,data string) {
	dataChannel <- data
	errorChannel <- ""
}

// WaitOnBothChannels will wait for the bytes and error and once received it will call corresponding bytes & error responders
func WaitOnBothChannels(c Client) {
	for i := 0; i < 2; i++ {
		select {
		case responseString := <-c.RespStringChannel:
			if responseString != "" {
				c.Success.Response(responseString)
			}
		case err := <-c.ErrStringChannel:
			// TODO handleErrorsBeforeSendingToClient(c, err)
			c.Failure.Error(err)
		}
	}
}