package chatbot

import (
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
