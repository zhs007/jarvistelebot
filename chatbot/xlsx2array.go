package chatbot

import (
	"bytes"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Xlsx2ArrayMap - xlsx to array, everything is string
func Xlsx2ArrayMap(data []byte) ([](map[string]interface{}), error) {
	r := bytes.NewReader(data)
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
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

				cl[fname[x]] = colCell
			}

			if hasdata {
				arr = append(arr, cl)
			}
		}
	}

	return arr, nil
}
