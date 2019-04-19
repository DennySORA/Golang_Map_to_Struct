package MapToStruct

import (
	"fmt"
	"reflect"
	"strings"
)

type decodeStruct struct {
	structData      interface{}
	mapToStructName map[string]string
	extend          map[string]interface{}
	key             string
	value           Values
}

func DecodeStructFromMap(mapData interface{}, structData interface{}, mapToStructName map[string]string, extend map[string]interface{}) error {
	mapType := checkType(mapData)
	if mapType == "__[]Map" {
		mapDataConvert := reflect.ValueOf(mapData)
		// -------------------------------------
		valElemType := reflect.ValueOf(structData).Elem().Type()
		DataCount := mapDataConvert.Len()
		NewSlice := reflect.MakeSlice(valElemType, DataCount, DataCount)
		// -------------------------------------
		for count := 0; count < mapDataConvert.Len(); count++ {
			currentField := NewSlice.Index(count)
			if mapToStructName == nil {
				mapToStructName = GetStructName(currentField)
			}
			DecodeStructFromMap(mapDataConvert.Index(count).Interface(), currentField, mapToStructName, extend)
		}
		reflect.ValueOf(structData).Elem().Set(NewSlice)
		return nil
	} else if mapType != "__Map" {
		return fmt.Errorf("Not is Map or []Map")
	} else {
		if mapToStructName == nil {
			mapToStructName = GetStructName(structData)
		}
		mapDataConvert := mapConvert(mapData)
		// -------------------------------------
		decodeStructInput := decodeStruct{
			structData:      structData,
			mapToStructName: mapToStructName,
			extend:          extend,
		}
		// -------------------------------------
		for key, value := range mapDataConvert {
			decodeStructInput.key = key
			decodeStructInput.value = value
			decodeStructInput.setField()
		}
		// -------------------------------------
		return nil
	}
}

func (std *decodeStruct) setField() error {
	// -------------------------------------
	var structValue reflect.Value
	if val, ok := std.structData.(reflect.Value); ok {
		structValue = val
	} else {
		structValue = reflect.ValueOf(std.structData).Elem()
	}
	// -------------------------------------
	key := std.mapToStructName[strings.ToLower(std.key)]
	// -------------------------------------
	structFieldValue := structValue.FieldByName(key)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", key)
	} else if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", key)
	}
	// -------------------------------------
	var val reflect.Value
	switch structFieldValue.Kind() {
	case reflect.Bool:
		val = reflect.ValueOf(std.value.GetBool())
	case reflect.Int:
		val = reflect.ValueOf(std.value.GetInt())
	case reflect.Int8:
		val = reflect.ValueOf(std.value.GetInt8())
	case reflect.Int16:
		val = reflect.ValueOf(std.value.GetInt16())
	case reflect.Int32:
		val = reflect.ValueOf(std.value.GetInt32())
	case reflect.Int64:
		val = reflect.ValueOf(std.value.GetInt64())
	case reflect.Uint:
		val = reflect.ValueOf(std.value.GetUInt())
	case reflect.Uint8:
		val = reflect.ValueOf(std.value.GetUInt8())
	case reflect.Uint16:
		val = reflect.ValueOf(std.value.GetUInt16())
	case reflect.Uint32:
		val = reflect.ValueOf(std.value.GetUInt32())
	case reflect.Uint64:
		val = reflect.ValueOf(std.value.GetUInt64())
	case reflect.Float32:
		val = reflect.ValueOf(std.value.GetFloat32())
	case reflect.Float64:
		val = reflect.ValueOf(std.value.GetFloat64())
	case reflect.String:
		val = reflect.ValueOf(std.value.GetString())
	}
	// -------------------------------------
	if everyList, ok := std.extend["Everys"]; ok {
		for _, funcEvery := range everyList.([]func(reflect.Value) reflect.Value) {
			val = funcEvery(val)
		}
	}
	if funcExtend, ok := std.extend[key]; ok {
		funcTemp := funcExtend.(func(reflect.Value) reflect.Value)
		val = funcTemp(val)
	}
	// -------------------------------------
	structFieldValue.Set(val)
	return nil
}

func checkType(checkData interface{}) string {
	kinds := reflect.ValueOf(checkData).Kind()
	if kinds == reflect.Map {
		return "__Map"
	} else if kinds == reflect.Slice {
		return "__[]Map"
	} else {
		return ""
	}
}

func GetStructName(obj interface{}) map[string]string {
	var structValue reflect.Value
	var ReturnMap map[string]string
	var ok bool

	if val, ok := obj.(reflect.Value); ok {
		structValue = val
	} else if objType := reflect.ValueOf(obj).Kind(); objType == reflect.Ptr {
		structValue = reflect.ValueOf(obj).Elem()
	} else {
		structValue = reflect.ValueOf(obj)
	}

	if objType := reflect.TypeOf(obj); objType.Kind() == reflect.Ptr {
		ReturnMap, ok = MapToSturctName[reflect.TypeOf(obj).Elem().Name()]
	} else {
		ReturnMap, ok = MapToSturctName[reflect.TypeOf(obj).Name()]
	}

	if !ok {
		ReturnMap = make(map[string]string)
	}

	for i := 0; i < structValue.NumField(); i++ {
		StuctInsideName := strings.ToLower(structValue.Type().Field(i).Name)
		if _, ok := ReturnMap[StuctInsideName]; !ok {
			ReturnMap[StuctInsideName] = structValue.Type().Field(i).Name
		}
	}
	return ReturnMap
}
