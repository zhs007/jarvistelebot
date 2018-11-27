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
