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
	"testing"
)

func TestFindPathVariables(t *testing.T) {
	findPathVariables("/api/{id}/book/{code}")
}

func TestRequest(t *testing.T) {
	//reqURL := "http://www.baidu.com"
	reqURL := "https://api.e-learn.io/watch?courseId=4666&lectureId=34947&coursewareId=59427"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bs, err := RequestClient(client, http.MethodGet, reqURL, nil, nil)

	fmt.Println("*************************************************")
	fmt.Println(string(bs))
	fmt.Println(err)
}
