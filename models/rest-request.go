package models

type RestRequest struct {
	Method      string
	Url         string
	Data        interface{}
	Headers     map[string]string
	Timeout     int
	QueryParams map[string]string
}

const defaultTimeout = 60

func NewRestRequest(method string, url string, data interface{},
	headers map[string]string, queryParams map[string]string, timeout int) *RestRequest {
	req := RestRequest{
		Method:      method,
		Url:         url,
		Data:        data,
		Headers:     headers,
		Timeout:     defaultTimeout,
		QueryParams: queryParams,
	}
	if timeout != 0 {
		req.Timeout = timeout
	}
	return &req
}
