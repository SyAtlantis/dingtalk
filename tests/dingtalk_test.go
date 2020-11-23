package tests

import (
	"github.com/zhaoyunxing92/dingtalk"
	"testing"
)

var dingTalk = dingtalk.NewDingTalk("dingdv9vbx9mcrj18g1n", "bDBJd67ct1Ik0GFqhNNWH4Lbo4aqZGglaE1wJ3mnbG6ANRjOruuGzs6Z0glNEU63")

func TestDingTalkGetToken(t *testing.T) {

	token, err := dingTalk.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}