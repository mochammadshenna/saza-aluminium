package state

var _httpHeaders httpHeaders

func init() {
	_httpHeaders = newHttpHeaders()
}

type httpHeader string

func (h httpHeader) String() string {
	return string(h)
}

type httpHeaders struct {
	Authorization httpHeader
	ContentType   httpHeader
	StartTime     httpHeader
	RequestId     httpHeader
}

func newHttpHeaders() httpHeaders {
	return httpHeaders{
		Authorization: "Authorization",
		ContentType:   "Content-Type",
		StartTime:     "Start-Time",
		RequestId:     "Request-Id",
	}
}

func HttpHeaders() httpHeaders {
	return _httpHeaders
}
