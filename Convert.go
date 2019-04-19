package MapToStruct

import (
	"reflect"
	"strconv"
)

// ===============================================================================

func mapConvert(mapData interface{}) map[string]Values {
	returnData := map[string]Values{}
	switch mapData.(type) {
	case map[string]string:
		for key, val := range mapData.(map[string]string) {
			returnData[key] = Parse(val)
		}
	case map[string]int:
		for key, val := range mapData.(map[string]int) {
			returnData[key] = Parse(val)
		}
	case map[string]int8:
		for key, val := range mapData.(map[string]int8) {
			returnData[key] = Parse(val)
		}
	case map[string]int16:
		for key, val := range mapData.(map[string]int16) {
			returnData[key] = Parse(val)
		}
	case map[string]int32:
		for key, val := range mapData.(map[string]int32) {
			returnData[key] = Parse(val)
		}
	case map[string]int64:
		for key, val := range mapData.(map[string]int64) {
			returnData[key] = Parse(val)
		}
	case map[string]uint:
		for key, val := range mapData.(map[string]uint) {
			returnData[key] = Parse(val)
		}
	case map[string]uint16:
		for key, val := range mapData.(map[string]uint16) {
			returnData[key] = Parse(val)
		}
	case map[string]uint32:
		for key, val := range mapData.(map[string]uint32) {
			returnData[key] = Parse(val)
		}
	case map[string]uint64:
		for key, val := range mapData.(map[string]uint64) {
			returnData[key] = Parse(val)
		}
	case map[string]float32:
		for key, val := range mapData.(map[string]float32) {
			returnData[key] = Parse(val)
		}
	case map[string]float64:
		for key, val := range mapData.(map[string]float64) {
			returnData[key] = Parse(val)
		}
	}
	return returnData
}

// ===============================================================================

type Values struct {
	Val  interface{}
	Kind reflect.Kind
}

// -----------------------------

func Parse(input interface{}) Values {
	return Values{Val: input, Kind: reflect.TypeOf(input).Kind()}
}

// -----------------------------

func (v *Values) GetBool() bool {
	switch v.Kind {
	case reflect.Bool:
		return v.Val.(bool)
	case reflect.Int:
		if v.Val.(int) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Int8:
		if v.Val.(int8) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Int16:
		if v.Val.(int16) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Int32:
		if v.Val.(int32) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Int64:
		if v.Val.(int64) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint:
		if v.Val.(uint) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint8:
		if v.Val.(uint8) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint16:
		if v.Val.(uint16) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint32:
		if v.Val.(uint32) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint64:
		if v.Val.(uint64) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Float32:
		if v.Val.(float32) != 0 {
			return true
		} else {
			return false
		}
	case reflect.Float64:
		if v.Val.(float64) != 0 {
			return true
		} else {
			return false
		}
	case reflect.String:
		if v.Val.(string) != "0" {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// -----------------------------

func (v *Values) GetInt8() int8 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return int8(v.Val.(int))
	case reflect.Int8:
		return int8(v.Val.(int8))
	case reflect.Int16:
		return int8(v.Val.(int16))
	case reflect.Int32:
		return int8(v.Val.(int32))
	case reflect.Int64:
		return int8(v.Val.(int64))
	case reflect.Uint:
		return int8(v.Val.(uint))
	case reflect.Uint8:
		return int8(v.Val.(uint8))
	case reflect.Uint16:
		return int8(v.Val.(uint16))
	case reflect.Uint32:
		return int8(v.Val.(uint32))
	case reflect.Uint64:
		return int8(v.Val.(uint64))
	case reflect.Float32:
		return int8(v.Val.(float32))
	case reflect.Float64:
		return int8(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return int8(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetInt16() int16 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return int16(v.Val.(int))
	case reflect.Int8:
		return int16(v.Val.(int8))
	case reflect.Int16:
		return int16(v.Val.(int16))
	case reflect.Int32:
		return int16(v.Val.(int32))
	case reflect.Int64:
		return int16(v.Val.(int64))
	case reflect.Uint:
		return int16(v.Val.(uint))
	case reflect.Uint8:
		return int16(v.Val.(uint8))
	case reflect.Uint16:
		return int16(v.Val.(uint16))
	case reflect.Uint32:
		return int16(v.Val.(uint32))
	case reflect.Uint64:
		return int16(v.Val.(uint64))
	case reflect.Float32:
		return int16(v.Val.(float32))
	case reflect.Float64:
		return int16(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return int16(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetInt32() int32 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return int32(v.Val.(int))
	case reflect.Int8:
		return int32(v.Val.(int8))
	case reflect.Int16:
		return int32(v.Val.(int16))
	case reflect.Int32:
		return int32(v.Val.(int32))
	case reflect.Int64:
		return int32(v.Val.(int64))
	case reflect.Uint:
		return int32(v.Val.(uint))
	case reflect.Uint8:
		return int32(v.Val.(uint8))
	case reflect.Uint16:
		return int32(v.Val.(uint16))
	case reflect.Uint32:
		return int32(v.Val.(uint32))
	case reflect.Uint64:
		return int32(v.Val.(uint64))
	case reflect.Float32:
		return int32(v.Val.(float32))
	case reflect.Float64:
		return int32(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return int32(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetInt64() int64 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return int64(v.Val.(int))
	case reflect.Int8:
		return int64(v.Val.(int8))
	case reflect.Int16:
		return int64(v.Val.(int16))
	case reflect.Int32:
		return int64(v.Val.(int32))
	case reflect.Int64:
		return int64(v.Val.(int64))
	case reflect.Uint:
		return int64(v.Val.(uint))
	case reflect.Uint8:
		return int64(v.Val.(uint8))
	case reflect.Uint16:
		return int64(v.Val.(uint16))
	case reflect.Uint32:
		return int64(v.Val.(uint32))
	case reflect.Uint64:
		return int64(v.Val.(uint64))
	case reflect.Float32:
		return int64(v.Val.(float32))
	case reflect.Float64:
		return int64(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return int64(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetUInt8() uint8 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return uint8(v.Val.(int))
	case reflect.Int8:
		return uint8(v.Val.(int8))
	case reflect.Int16:
		return uint8(v.Val.(int16))
	case reflect.Int32:
		return uint8(v.Val.(int32))
	case reflect.Int64:
		return uint8(v.Val.(int64))
	case reflect.Uint:
		return uint8(v.Val.(uint))
	case reflect.Uint8:
		return uint8(v.Val.(uint8))
	case reflect.Uint16:
		return uint8(v.Val.(uint16))
	case reflect.Uint32:
		return uint8(v.Val.(uint32))
	case reflect.Uint64:
		return uint8(v.Val.(uint64))
	case reflect.Float32:
		return uint8(v.Val.(float32))
	case reflect.Float64:
		return uint8(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return uint8(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetUInt16() uint16 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return uint16(v.Val.(int))
	case reflect.Int8:
		return uint16(v.Val.(int8))
	case reflect.Int16:
		return uint16(v.Val.(int16))
	case reflect.Int32:
		return uint16(v.Val.(int32))
	case reflect.Int64:
		return uint16(v.Val.(int64))
	case reflect.Uint:
		return uint16(v.Val.(uint))
	case reflect.Uint8:
		return uint16(v.Val.(uint8))
	case reflect.Uint16:
		return uint16(v.Val.(uint16))
	case reflect.Uint32:
		return uint16(v.Val.(uint32))
	case reflect.Uint64:
		return uint16(v.Val.(uint64))
	case reflect.Float32:
		return uint16(v.Val.(float32))
	case reflect.Float64:
		return uint16(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return uint16(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetUInt32() uint32 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return uint32(v.Val.(int))
	case reflect.Int8:
		return uint32(v.Val.(int8))
	case reflect.Int16:
		return uint32(v.Val.(int16))
	case reflect.Int32:
		return uint32(v.Val.(int32))
	case reflect.Int64:
		return uint32(v.Val.(int64))
	case reflect.Uint:
		return uint32(v.Val.(uint))
	case reflect.Uint8:
		return uint32(v.Val.(uint8))
	case reflect.Uint16:
		return uint32(v.Val.(uint16))
	case reflect.Uint32:
		return uint32(v.Val.(uint32))
	case reflect.Uint64:
		return uint32(v.Val.(uint64))
	case reflect.Float32:
		return uint32(v.Val.(float32))
	case reflect.Float64:
		return uint32(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return uint32(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetUInt64() uint64 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return uint64(v.Val.(int))
	case reflect.Int8:
		return uint64(v.Val.(int8))
	case reflect.Int16:
		return uint64(v.Val.(int16))
	case reflect.Int32:
		return uint64(v.Val.(int32))
	case reflect.Int64:
		return uint64(v.Val.(int64))
	case reflect.Uint:
		return uint64(v.Val.(uint))
	case reflect.Uint8:
		return uint64(v.Val.(uint8))
	case reflect.Uint16:
		return uint64(v.Val.(uint16))
	case reflect.Uint32:
		return uint64(v.Val.(uint32))
	case reflect.Uint64:
		return uint64(v.Val.(uint64))
	case reflect.Float32:
		return uint64(v.Val.(float32))
	case reflect.Float64:
		return uint64(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return uint64(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetInt() int {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return int(v.Val.(int))
	case reflect.Int8:
		return int(v.Val.(int8))
	case reflect.Int16:
		return int(v.Val.(int16))
	case reflect.Int32:
		return int(v.Val.(int32))
	case reflect.Int64:
		return int(v.Val.(int64))
	case reflect.Uint:
		return int(v.Val.(uint))
	case reflect.Uint8:
		return int(v.Val.(uint8))
	case reflect.Uint16:
		return int(v.Val.(uint16))
	case reflect.Uint32:
		return int(v.Val.(uint32))
	case reflect.Uint64:
		return int(v.Val.(uint64))
	case reflect.Float32:
		return int(v.Val.(float32))
	case reflect.Float64:
		return int(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return int(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetUInt() uint {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return uint(v.Val.(int))
	case reflect.Int8:
		return uint(v.Val.(int8))
	case reflect.Int16:
		return uint(v.Val.(int16))
	case reflect.Int32:
		return uint(v.Val.(int32))
	case reflect.Int64:
		return uint(v.Val.(int64))
	case reflect.Uint:
		return uint(v.Val.(uint))
	case reflect.Uint8:
		return uint(v.Val.(uint8))
	case reflect.Uint16:
		return uint(v.Val.(uint16))
	case reflect.Uint32:
		return uint(v.Val.(uint32))
	case reflect.Uint64:
		return uint(v.Val.(uint64))
	case reflect.Float32:
		return uint(v.Val.(float32))
	case reflect.Float64:
		return uint(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return uint(number)
		}
	default:
		return 0
	}
}

// -----------------------------

func (v *Values) GetFloat32() float32 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return float32(v.Val.(int))
	case reflect.Int8:
		return float32(v.Val.(int8))
	case reflect.Int16:
		return float32(v.Val.(int16))
	case reflect.Int32:
		return float32(v.Val.(int32))
	case reflect.Int64:
		return float32(v.Val.(int64))
	case reflect.Uint:
		return float32(v.Val.(uint))
	case reflect.Uint8:
		return float32(v.Val.(uint8))
	case reflect.Uint16:
		return float32(v.Val.(uint16))
	case reflect.Uint32:
		return float32(v.Val.(uint32))
	case reflect.Uint64:
		return float32(v.Val.(uint64))
	case reflect.Float32:
		return float32(v.Val.(float32))
	case reflect.Float64:
		return float32(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return float32(number)
		}
	default:
		return 0
	}
}

func (v *Values) GetFloat64() float64 {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return 1
		} else {
			return 0
		}
	case reflect.Int:
		return float64(v.Val.(int))
	case reflect.Int8:
		return float64(v.Val.(int8))
	case reflect.Int16:
		return float64(v.Val.(int16))
	case reflect.Int32:
		return float64(v.Val.(int32))
	case reflect.Int64:
		return float64(v.Val.(int64))
	case reflect.Uint:
		return float64(v.Val.(uint))
	case reflect.Uint8:
		return float64(v.Val.(uint8))
	case reflect.Uint16:
		return float64(v.Val.(uint16))
	case reflect.Uint32:
		return float64(v.Val.(uint32))
	case reflect.Uint64:
		return float64(v.Val.(uint64))
	case reflect.Float32:
		return float64(v.Val.(float32))
	case reflect.Float64:
		return float64(v.Val.(float64))
	case reflect.String:
		number, err := strconv.ParseFloat(v.Val.(string), 64)
		if err != nil {
			return 0
		} else {
			return float64(number)
		}
	default:
		return 0
	}
}

// -----------------------------

func (v *Values) GetString() string {
	switch v.Kind {
	case reflect.Bool:
		if v.Val.(bool) {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int:
		return strconv.FormatInt(int64(v.Val.(int)), 10)
	case reflect.Int8:
		return strconv.FormatInt(int64(v.Val.(int8)), 10)
	case reflect.Int16:
		return strconv.FormatInt(int64(v.Val.(int16)), 10)
	case reflect.Int32:
		return strconv.FormatInt(int64(v.Val.(int32)), 10)
	case reflect.Int64:
		return strconv.FormatInt(int64(v.Val.(int64)), 10)
	case reflect.Uint:
		return strconv.FormatUint(uint64(v.Val.(uint)), 10)
	case reflect.Uint8:
		return strconv.FormatUint(uint64(v.Val.(uint8)), 10)
	case reflect.Uint16:
		return strconv.FormatUint(uint64(v.Val.(uint16)), 10)
	case reflect.Uint32:
		return strconv.FormatUint(uint64(v.Val.(uint32)), 10)
	case reflect.Uint64:
		return strconv.FormatUint(uint64(v.Val.(uint64)), 10)
	case reflect.Float32:
		return strconv.FormatFloat(float64(v.Val.(float32)), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(float64(v.Val.(float32)), 'f', -1, 64)
	case reflect.String:
		return v.Val.(string)
	default:
		return ""
	}
}
