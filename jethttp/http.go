package jethttp

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpClient(scheme string) *http.Client {
	if scheme == "https" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		return client
	}

	return &http.Client{}
}
func HttpResponse(client *http.Client, myurl string) (*http.Response, error) {
	resp, err := client.Get(myurl)
	defer resp.Body.Close()
	return resp, err

}

func HttpResponseBody(client *http.Client, myurl string) (body []byte, e error) {
	resp, err := HttpResponse(client, myurl)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func HttpPostUrl(myurl string) {

}

func HttpRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
