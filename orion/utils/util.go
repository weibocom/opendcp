package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(v interface{}) (int, error) {
	switch v.(type) {
	case int:
		return v.(int), nil
	case float64:
		return int(v.(float64)), nil
	case string:
		i, err := strconv.Atoi(v.(string))
		if err == nil {
			return i, nil
		}
	}

	return -1, errors.New("cannot convert to int:" + fmt.Sprintln(v) +
		", type : " + reflect.TypeOf(v).String())
}
