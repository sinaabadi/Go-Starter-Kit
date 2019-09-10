package utils

import (
	"github.com/go-resty/resty"
	"starter-kit/models"
	"time"
)

func init() {
}
func NewClient(timeout time.Duration) *resty.Client {
	return resty.New().
		SetTimeout(timeout)

}
func Call(req *models.RestRequest, resultModel *interface{}) error {
	RestClient := NewClient(time.Duration(req.Timeout * int(time.Second)))
	var err error
	switch req.Method {
	case `GET`:
		_, err = RestClient.
			R().
			SetQueryParams(req.QueryParams).
			SetResult(resultModel).
			SetHeaders(req.Headers).
			Get(req.Url)
	case `POST`:
		_, err = RestClient.
			R().
			SetBody(req.Data).
			SetResult(resultModel).
			SetQueryParams(req.QueryParams).
			SetHeaders(req.Headers).
			Post(req.Url)
	case `PUT`:
		_, err = RestClient.
			R().
			SetBody(req.Data).
			SetQueryParams(req.QueryParams).
			SetResult(resultModel).
			SetHeaders(req.Headers).
			Put(req.Url)
	case `DELETE`:
		_, err = RestClient.
			R().
			SetBody(req.Data).
			SetResult(resultModel).
			SetQueryParams(req.QueryParams).
			SetHeaders(req.Headers).
			Delete(req.Url)
	}
	return err
}
