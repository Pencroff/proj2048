package general

import (
	"reflect"
)

// InterfaceSlice transform any slice to []interface{} slice
func InterfaceSlice(slice interface{}) []interface{} {
	switch slice := slice.(type) {
	case nil:
		return nil
	case []string:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []bool:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []int:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []int8:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []int16:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []int32:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []int64:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []uint:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []uint8:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []uint16:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []uint32:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []uint64:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []float32:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	case []float64:
		res := make([]interface{}, len(slice))
		for i, v := range slice {
			res[i] = v
		}
		return res
	default:
		s := reflect.ValueOf(slice)
		if s.Kind() != reflect.Slice {
			panic("InterfaceSlice() given a non-slice type")
		}

		res := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Interface()
		}

		return res
	}

}
