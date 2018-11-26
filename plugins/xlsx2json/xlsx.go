package pluginxlsx2json

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func toJSON(data []byte) (string, error) {
	r := bytes.NewReader(data)
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return "", err
	}

	fname := make(map[int]string)
	var arr []map[string]interface{}
	mapname := xlsx.GetSheetMap()
	// sname := xlsx.GetSheetName(0)
	// fmt.Printf("sname-%v", mapname)
	for y, row := range xlsx.GetRows(mapname[1]) {
		// fmt.Printf("y-%v", y)

		if y == 0 {
			for x, colCell := range row {
				fname[x] = colCell
			}

			// fmt.Printf("%v", fname)
		} else {
			cl := make(map[string]interface{})
			hasdata := false
			for x, colCell := range row {
				if colCell == "" {
					continue
				}

				hasdata = true

				i, err := strconv.ParseInt(colCell, 10, 64)
				if err == nil {
					cl[fname[x]] = i

					continue
				}

				f, err := strconv.ParseFloat(colCell, 64)
				if err == nil {
					cl[fname[x]] = f

					continue
				}

				b, err := strconv.ParseBool(colCell)
				if err == nil {
					cl[fname[x]] = b

					continue
				}

				cl[fname[x]] = colCell
			}

			if hasdata {
				arr = append(arr, cl)
			}
		}
	}

	jsonStr, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

func toJSONMap(data []byte, idname string) (string, error) {
	r := bytes.NewReader(data)
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return "", err
	}

	fname := make(map[int]string)
	// var obj map[string](map[string]interface{})
	var arr [](map[string]interface{})
	mapname := xlsx.GetSheetMap()
	idnameisint := true
	// sname := xlsx.GetSheetName(0)
	// fmt.Printf("sname-%v", mapname)
	for y, row := range xlsx.GetRows(mapname[1]) {
		// fmt.Printf("y-%v", y)

		if y == 0 {
			for x, colCell := range row {
				fname[x] = colCell
			}

			// fmt.Printf("%v", fname)
		} else {
			cl := make(map[string]interface{})
			hasdata := false
			for x, colCell := range row {
				if colCell == "" {
					continue
				}

				hasdata = true

				i, err := strconv.ParseInt(colCell, 10, 64)
				if err == nil {
					cl[fname[x]] = i

					continue
				}

				if fname[x] == idname {
					idnameisint = false
				}

				f, err := strconv.ParseFloat(colCell, 64)
				if err == nil {
					cl[fname[x]] = f

					continue
				}

				b, err := strconv.ParseBool(colCell)
				if err == nil {
					cl[fname[x]] = b

					continue
				}

				cl[fname[x]] = colCell
			}

			if hasdata {
				arr = append(arr, cl)
				// obj[cl[idname].string] = cl
			}
		}
	}

	if idnameisint {
		obj := make(map[int64](map[string]interface{}))

		for _, v := range arr {
			idnameint, ok := v[idname].(int64)
			if ok {
				_, ok = obj[idnameint]
				if ok {
					return "", ErrJSONObjSameKey
				}

				obj[idnameint] = v
			} else {
				return "", ErrJSONObjIDKeyNotInt
			}
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return "", err
		}

		return string(jsonStr), nil
	}

	obj := make(map[string](map[string]interface{}))

	for _, v := range arr {
		idnamestr, ok := v[idname].(string)
		if ok {
			_, ok = obj[idnamestr]
			if ok {
				return "", ErrJSONObjSameKey
			}

			obj[idnamestr] = v
		} else {
			return "", ErrJSONObjIDKeyNotString
		}
	}

	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}
