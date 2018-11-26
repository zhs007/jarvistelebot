package pluginxlsx2json

import (
	"io/ioutil"
	"testing"
)

func Test_toJSON(t *testing.T) {
	dat, err := ioutil.ReadFile("../../test/list.xlsx")
	if err != nil {
		t.Fatalf("Test_toJSON ReadFile Err %v", err)
	}

	str, err := toJSON(dat)
	if err != nil {
		t.Fatalf("Test_toJSON toJSON Err %v", err)
	}

	dat, err = ioutil.ReadFile("../../test/list1.xlsx")
	if err != nil {
		t.Fatalf("Test_toJSON ReadFile Err %v", err)
	}

	str, err = toJSON(dat)
	if err != nil {
		t.Fatalf("Test_toJSON toJSON Err %v", err)
	}

	t.Logf("Test_toJSON OK %v", str)
}

func Test_toJSONMap(t *testing.T) {
	dat, err := ioutil.ReadFile("../../test/map.xlsx")
	if err != nil {
		t.Fatalf("Test_toJSONMap ReadFile Err %v", err)
	}

	str, err := toJSONMap(dat, "name")
	if err != nil {
		t.Fatalf("Test_toJSONMap toJSON Err %v", err)
	}

	t.Logf("Test_toJSONMap OK %v", str)
}
