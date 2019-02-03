package network

//Timeout Constants
const (
	DEFAULT_NETCLIENT_TIMEOUT               int = 30
)

//Network Constants
const DEFAULT_RETRY int = 0

//Log Constants
const LOG_NETWORK_TAG string = "NetworkLayer-- "

//Error Constants
const (
	NETWORK_REQUEST_ERROR       int = 10001
	NETWORK_RESPONSE_ERROR      int = 10002
	NETWORK_SERIALIZATION_ERROR int = 10003
	NETWORK_ERROR               int = 10004 // GENERIC NETWORK ERROR
	NETWORK_PANIC_ERROR         int = 10005
)
