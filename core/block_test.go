package core

import (
	"testing"
)

func Test(t *testing.T) {
	godBlock := createGodBlock("创世纪块");
	secondBlock := CreateBlock(godBlock);
	secondBlock.Body = "我是第二个块";
	if(secondBlock.Body == "我是第二个块"){
		t.Log("ok")
	} else {
		t.Fail();
	}
}