// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-06-03 15:20
// version: 1.0.0
// desc   :

package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestFindPathVariables(t *testing.T) {
	findPathVariables("/api/{id}/book/{code}")
}

func TestRequest(t *testing.T) {
	//reqURL := "http://www.baidu.com"
	reqURL := "https://api.e-learn.io/watch?courseId=4666&lectureId=34947&coursewareId=59427"

	var PTransport = &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: PTransport}

	os.Setenv("HTTPS_PROXY", "https://222.74.202.245:8080")

	bs, err := RequestClient(client, http.MethodGet, reqURL, nil, nil)

	fmt.Println(string(bs))
	fmt.Println(err)
}
