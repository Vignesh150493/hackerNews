package network

import "fmt"

type NetworkError struct {
	statusCode int
	Body       string
}

func (n NetworkError) IsAuthError() bool {
	return n.statusCode == 401
}

func (n *NetworkError) IsResourceNotFound() bool {
	return n.statusCode == 404
}

func (n *NetworkError) IsServerError() bool {
	return n.statusCode >= 500
}

func (n *NetworkError) IsAccessDenied() bool {
	return n.statusCode == 403
}

func (n NetworkError) Error() string {
	return fmt.Sprintf("Network error with status code: %d", n.statusCode)
}

// InvalidUrlError , this error occurs when the url passed for the network request is invalid. ex: (does not contain protocol etc.)
type InvalidUrlError struct {
	url string
}

func (e *InvalidUrlError) Error() string {
	return fmt.Sprintf("The URL %s is not a valid url", e.url)
}
