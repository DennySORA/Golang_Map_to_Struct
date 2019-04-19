package main

import (
	MapToStruct "MTS"
	"fmt"
	"reflect"
	"regexp"
)

type A struct {
	Data float32
}

func main() {
	InA := map[string]int{
		"Data": 10,
	}
	OutA := &A{}
	err := MapToStruct.DecodeStructFromMap(InA, OutA, nil, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(OutA)
	}

	// =====================================

	InB := []map[string]string{
		map[string]string{
			"Data": "10",
		},
		map[string]string{
			"Data": "15",
		},
	}
	OutB := &[]A{}
	err = MapToStruct.DecodeStructFromMap(InB, OutB, nil, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(OutB)
	}

	// =====================================

	InC := []map[string]string{
		map[string]string{
			"Data": "10.01520",
		},
		map[string]string{
			"Data": "15.0100",
		},
	}
	OutC := &[]A{}
	InFunc := map[string]interface{}{
		"Data": ConvertStrCut,
	}
	err = MapToStruct.DecodeStructFromMap(InC, OutC, nil, InFunc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(OutC)
	}
	// =====================================

	InD := []map[string]string{
		map[string]string{
			"app": "10.01520",
		},
		map[string]string{
			"app": "15.0100",
		},
	}
	OutD := &[]A{}
	InFunc = map[string]interface{}{
		"Data": ConvertStrCut,
	}

	NameA := MapToStruct.GetStructName(A{})
	fmt.Println(NameA)

	err = MapToStruct.DecodeStructFromMap(InD, OutD, NameA, InFunc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(OutD)
	}

	// =====================================

	NameB := MapToStruct.GetStructName(&A{})
	fmt.Println(NameB)
}

func ConvertStrCut(number reflect.Value) reflect.Value {
	r, _ := regexp.Compile(`(-?[0-9]+)((\.0+)|(\.[0-9]+?))0*$`)
	if val, ok := number.Interface().(string); ok {
		Temp := r.ReplaceAllString(val, "$1$4")
		return reflect.ValueOf(Temp)
	} else {
		return number
	}
}
