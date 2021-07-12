// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-06-03 15:20
// version: 1.0.0
// desc   :

package client

import "testing"

func TestFindPathVariables(t *testing.T) {
	findPathVariables("/api/{id}/book/{code}")
}
