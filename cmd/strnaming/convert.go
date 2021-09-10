// Copyright (c) 2021 startdusk. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"

	"github.com/tidwall/gjson"
)

// Converter for strnaming instance
type Converter interface {
	Convert(str string) string
}

func convert(converter Converter, input string) ([]byte, error) {
	if !gjson.Valid(input) {
		return nil, errors.New("invalid json")
	}
	res := gjson.Parse(input)

	var deepConvert func(res gjson.Result) interface{}
	deepConvert = func(res gjson.Result) interface{} {
		isObj, isArr := res.IsObject(), res.IsArray()
		if !isArr && !isObj {
			return caseType(res)
		}
		if isObj {
			dest := make(map[string]interface{})
			res.ForEach(func(key, value gjson.Result) bool {
				newKey := converter.Convert(key.String())
				dest[newKey] = deepConvert(value)
				return true
			})
			return dest
		} else if isArr {
			var dest []interface{}
			res.ForEach(func(_, value gjson.Result) bool {
				dest = append(dest, deepConvert(value))
				return true
			})
			return dest
		}

		return nil
	}

	destObj := deepConvert(res)
	return json.MarshalIndent(destObj, "", "    ")
}

func caseType(t gjson.Result) interface{} {
	switch t.Type {
	default:
		return ""
	case gjson.Null:
		return ""
	case gjson.False:
		return false
	case gjson.Number:
		return t.Int()
	case gjson.String:
		return t.String()
	case gjson.True:
		return true
	}
}
