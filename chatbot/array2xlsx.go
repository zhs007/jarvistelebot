package chatbot

import (
	"bufio"
	"bytes"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// getCellName - get cell name with x, y
func getCellName(x int, y int) string {
	if x < 0 || y < 0 {
		return ""
	}

	if x >= 27*26 {
		return ""
	}

	x0 := x / 26
	x1 := x % 26

	cn := ""
	a := 'A'

	if x0 == 0 {
		cn += string(a + rune(x1))
	} else {
		cn += string(a + rune(x0-1))
		cn += string(a + rune(x1))
	}

	cn += strconv.Itoa(y + 1)

	return cn
}

// getObjName - get object name
func getObjName(obj interface{}) []string {
	t := reflect.TypeOf(obj)

	var lst []string
	for i := 0; i < t.NumField(); i++ {
		lst = append(lst, t.Field(i).Name)
	}

	return lst
}

// getKind - get kind
func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

// obj2map - object to map
func obj2map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		dataKind := getKind(v.Field(i))
		switch dataKind {
		case reflect.Bool:
			data[t.Field(i).Name] = v.Field(i).Bool()
		case reflect.String:
			data[t.Field(i).Name] = v.Field(i).String()
		case reflect.Int:
			data[t.Field(i).Name] = v.Field(i).Int()
		case reflect.Uint:
			data[t.Field(i).Name] = v.Field(i).Uint()
		case reflect.Float32:
			data[t.Field(i).Name] = v.Field(i).Float()
		}
	}

	return data
}

// Array2xlsx - array to xlsx
func Array2xlsx(arr []interface{}) ([]byte, error) {
	if len(arr) <= 0 {
		return nil, ErrEmptyArray
	}

	xlsx := excelize.NewFile()
	lsthead := getObjName(arr[0])

	for x, v := range lsthead {
		xlsx.SetCellValue("Sheet1", getCellName(x, 0), v)
	}

	for y, v := range arr {
		m := obj2map(v)

		for x, hv := range lsthead {
			cv, ok := m[hv]
			if ok {
				xlsx.SetCellValue("Sheet1", getCellName(x, y+1), cv)
			}
		}
	}

	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	err := xlsx.Write(w)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
