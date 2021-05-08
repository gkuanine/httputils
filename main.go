package httputils

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func NewRequest() *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
		cookies: map[string]string{},
	}
	return r
}

// Debug model
func Debug(v bool) *Request {
	r := NewRequest()
	return r.Debug(v)
}

func Jar(v http.CookieJar) *Request {
	r := NewRequest()
	return r.Jar(v)
}

func DisableKeepAlives(v bool) *Request {
	r := NewRequest()
	return r.DisableKeepAlives(v)
}

func CheckRedirect(v func(req *http.Request, via []*http.Request) error) *Request {
	r := NewRequest()
	return r.CheckRedirect(v)
}

func TLSClient(v *tls.Config) *Request {
	r := NewRequest()
	return r.SetTLSClient(v)
}

func SetTLSClient(v *tls.Config) *Request {
	r := NewRequest()
	return r.SetTLSClient(v)
}

func SetHeaders(headers map[string]string) *Request {
	r := NewRequest()
	return r.SetHeaders(headers)
}

func SetCookies(cookies map[string]string) *Request {
	r := NewRequest()
	return r.SetCookies(cookies)
}

func SetBasicAuth(username, password string) *Request {
	r := NewRequest()
	return r.SetBasicAuth(username, password)
}

func JSON() *Request {
	r := NewRequest()
	return r.JSON()
}

func Proxy(v func(*http.Request) (*url.URL, error)) *Request {
	r := NewRequest()
	return r.Proxy(v)
}

func SetTimeout(d time.Duration) *Request {
	r := NewRequest()
	return r.SetTimeout(d)
}

func Transport(v *http.Transport) *Request {
	r := NewRequest()
	return r.Transport(v)
}

// Get is a get http request
func Get(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Get(url, data...)
}

// Get is a get http request
func GetRetry(url string, retry int, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.GetRetry(url, retry, data...)
}

// Post is a post http request
func PostRetry(url string, retry int, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.PostRetry(url, retry, data...)
}

func Post(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Post(url, data...)
}

// Post is a post http request
func PostJsonRetry(url string, retry int, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.PostJsonRetry(url, retry, data...)
}

// Put is a put http request
func Put(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Put(url, data...)
}

// Delete is a delete http request
func Delete(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Delete(url, data...)
}

// Upload file
func Upload(url, filename, fileinput string) (*Response, error) {
	r := NewRequest()
	return r.Upload(url, filename, fileinput)
}
