package chatbot

import (
	"strconv"
	"testing"
)

func Test_getCellName(t *testing.T) {

	type testdata struct {
		cellname string
		x        int
		y        int
	}

	arrok := []testdata{
		testdata{cellname: "A1", x: 0, y: 0},
		testdata{cellname: "A2", x: 0, y: 1},
		testdata{cellname: "B2", x: 1, y: 1},
		testdata{cellname: "Z2", x: 25, y: 1},
		testdata{cellname: "AA11", x: 26, y: 10},
		testdata{cellname: "BB11", x: 53, y: 10},
		testdata{cellname: "ZC11", x: 26 + 25*26 + 2, y: 10},
	}

	for _, v := range arrok {
		cr := getCellName(v.x, v.y)
		if cr != v.cellname {
			t.Fatalf("Test_getCellName Err x:%v y:%v cellname:%v ret:%v", v.x, v.y, v.cellname, cr)
		}
	}

	t.Logf("Test_getCellName OK")
}

func Test_obj2map(t *testing.T) {

	type testdata struct {
		a int
		b bool
		c int32
		d int64
		e uint
		f uint8
		g uint16
		h uint32
		i uint64
		j int8
		k int16
		l string
		m float32
		n float64
	}

	lst := obj2map(testdata{a: 1, b: false, c: 2, d: 3, e: 4, f: 5, g: 6, h: 7, i: 8, j: 9, k: 10, l: "11", m: 1.2, n: 1.3})
	if len(lst) != 14 {
		t.Fatalf("Test_obj2map Err %v", len(lst))
	}

	t.Logf("Test_obj2map OK")
}

func Test_getMapObjName(t *testing.T) {

	mapobj0 := make(map[string]interface{})
	mapobj0["a"] = 1
	mapobj0["a"+strconv.Itoa(1)] = 2
	mapobj0["a"+strconv.Itoa(2)] = 3

	mapobj1 := make(map[string]interface{})
	mapobj1["a"] = 1
	mapobj1["a"+strconv.Itoa(1)] = 2
	mapobj1["b"] = 4
	mapobj1["a"+strconv.Itoa(3)] = 5

	var arr [](map[string]interface{})

	arr = append(arr, mapobj0)
	arr = append(arr, mapobj1)

	lst := getMapObjName(arr)
	if len(lst) != 5 {
		t.Fatalf("Test_getMapObjName Err %v", len(lst))
	}

	t.Logf("Test_getMapObjName OK")
}
