// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-01-26 11:01
// version: 1.0.0
// desc   :

package godis

import (
	"testing"
)

type TG struct {
	ID   int
	Name string
}

func TestGob(t *testing.T) {
	tg := TG{
		ID:   1,
		Name: "AA",
	}
	gb, err := toGob(tg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(gb)

	var tgs []TG

	err = gobStringToInterfaceSlice(&tgs, []string{gb}...)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tgs)
}
