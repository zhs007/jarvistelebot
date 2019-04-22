package chatbot

import (
	"testing"
)

func Test_GetFileNameFromFullPathNoExt(t *testing.T) {
	type strarr struct {
		src  string
		dest string
	}

	arrok := []strarr{
		strarr{src: "../../test/list.xlsx", dest: "list"},
		strarr{src: "list.xlsx", dest: "list"},
		strarr{src: "list.1.xlsx", dest: "list.1"},
		strarr{src: "list-1.xlsx", dest: "list-1"},
		strarr{src: "list 1.xlsx", dest: "list 1"},
	}

	for _, v := range arrok {
		cr := GetFileNameFromFullPathNoExt(v.src)
		if cr != v.dest {
			t.Fatalf("Test_GetFileNameFromFullPathNoExt Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
		}
	}

	t.Logf("Test_GetFileNameFromFullPathNoExt OK")
}

func Test_GetFileNameFromFullPath(t *testing.T) {
	type strarr struct {
		src  string
		dest string
	}

	arrok := []strarr{
		strarr{src: "../../test/list.xlsx", dest: "list.xlsx"},
		strarr{src: "./list.xlsx", dest: "list.xlsx"},
		strarr{src: "list.xlsx", dest: "list.xlsx"},
		strarr{src: "list.1.xlsx", dest: "list.1.xlsx"},
		strarr{src: "list-1.xlsx", dest: "list-1.xlsx"},
		strarr{src: "list 1.xlsx", dest: "list 1.xlsx"},
	}

	for _, v := range arrok {
		cr := GetFileNameFromFullPath(v.src)
		if cr != v.dest {
			t.Fatalf("Test_GetFileNameFromFullPath Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
		}
	}

	t.Logf("Test_GetFileNameFromFullPath OK")
}

func Test_SplitString(t *testing.T) {
	type strarr struct {
		src  string
		dest []string
	}

	arrok := []strarr{
		strarr{src: "getdtdata -m gamedatareport -s 2019-04-17 -e 2019-04-17", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: "getdtdata", dest: []string{"getdtdata"}},
		strarr{src: "   getdtdata", dest: []string{"getdtdata"}},
		strarr{src: "   getdtdata   ", dest: []string{"getdtdata"}},
		strarr{src: "getdtdata   ", dest: []string{"getdtdata"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    2019-04-17   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    \"2019-04-17\"   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    \"2019-04-17   \"   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17   ", "-e", "2019-04-17"}},
	}

	for _, v := range arrok {
		cr := SplitString(v.src)
		if len(cr) != len(v.dest) {
			t.Fatalf("Test_SplitString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
		}

		for i, sv := range cr {
			if sv != v.dest[i] {
				t.Fatalf("Test_SplitString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
			}
		}
	}

	t.Logf("Test_SplitString OK")
}
