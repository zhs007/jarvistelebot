package pluginxlsx2json

import "errors"

var (
	// ErrJSONObjIDKeyNotInt - some key is not int
	ErrJSONObjIDKeyNotInt = errors.New("some key is not int")
	// ErrJSONObjIDKeyNotString - some key is not string
	ErrJSONObjIDKeyNotString = errors.New("some key is not string")
	// ErrJSONObjSameKey - JSON has same key
	ErrJSONObjSameKey = errors.New("JSON has same key")
)
