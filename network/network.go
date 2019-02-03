package network

import (
	"io/ioutil"
	"net/http"
	"time"
)

var HTTPClient NetworkClient

type TimeOutOptions struct {
	NetClientTimeout             int // net Http data timeout
}

type NetworkClient struct {
	netClient *http.Client
}

type NetworkResponse struct {
	Resp string
	Err  error
}

func init() {
	HTTPClient = createNetworkClient(createDefaultNetClientTimoutOpts())
}

func createNetworkClient(timeOutOptions TimeOutOptions) NetworkClient {
	var netClient = NetworkClient{
		netClient: &http.Client{
			Timeout: time.Duration(timeOutOptions.NetClientTimeout) * time.Second,
		},
	}
	return netClient
}

func (client NetworkClient) GETRequest(path string, networkResult chan NetworkResponse) {
	resp, err := client.netClient.Get(path)
	if err != nil {
		//Handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Handle Serialization errors
	} else {
		networkResult <- NetworkResponse{Resp: string(body),}
	}
}

//PRIVATE METHODS
//DefaultOptions - configurable
func createDefaultNetClientTimoutOpts() TimeOutOptions {
	var timeoutOptions = TimeOutOptions{
		NetClientTimeout: DEFAULT_NETCLIENT_TIMEOUT,
	}
	return timeoutOptions
}
