package chatbot

import (
	"io/ioutil"
	"testing"
)

func Test_Xlsx2ArrayMap(t *testing.T) {
	dat, err := ioutil.ReadFile("../test/list.xlsx")
	if err != nil {
		t.Fatalf("Test_Xlsx2ArrayMap ReadFile Err %v", err)
	}

	str, err := Xlsx2ArrayMap(dat)
	if err != nil {
		t.Fatalf("Test_Xlsx2ArrayMap toJSON Err %v", err)
	}

	dat, err = ioutil.ReadFile("../test/list1.xlsx")
	if err != nil {
		t.Fatalf("Test_Xlsx2ArrayMap ReadFile Err %v", err)
	}

	str, err = Xlsx2ArrayMap(dat)
	if err != nil {
		t.Fatalf("Test_Xlsx2ArrayMap toJSON Err %v", err)
	}

	t.Logf("Test_Xlsx2ArrayMap OK %v", str)
}
