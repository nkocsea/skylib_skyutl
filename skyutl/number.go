package skyutl

import (
	"fmt"
	"strconv"
)

//ToInt64 convert interface{} to int64.
func ToInt64(source interface{}) (int64, error) {
	if source == nil {
		return 0, nil
	}

	str := fmt.Sprintf("%v", source)

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//ToFloat64 convert interface{} to float64.
func ToFloat64(source interface{}) (float64, error) {
	if source == nil {
		return 0, nil
	}

	str := fmt.Sprintf("%v", source)

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//ToF64 convert interface{} to float64.
func ToF64(source interface{}) float64 {
	num, err := ToFloat64(source)
	if err != nil {
		return 0.0
	}

	return num
}

//ToF64Ptr convert interface{} to *float64.
func ToF64Ptr(source interface{}) *float64 {
	num, err := ToFloat64(source)
	if err != nil {
		return nil
	}

	return &num
}

//ToI64 convert interface{} to int64.
func ToI64(source interface{}) int64 {
	num, err := ToInt64(source)
	if err != nil {
		return 0
	}

	return num
}

//ToI64Ptr convert interface{} to *int64.
func ToI64Ptr(source interface{}) *int64 {
	num, err := ToInt64(source)
	if err != nil {
		return nil
	}

	return &num
}

//ToInt32 convert interface{} to int32.
func ToInt32(source interface{}) (int32, error) {
	num, err := ToInt64(source)
	if err != nil {
		return 0, err
	}

	return int32(num), nil
}

//ToI32 convert interface{} to int32.
func ToI32(source interface{}) int32 {
	num, err := ToInt32(source)
	if err != nil {
		return 0
	}

	return num
}

//ToI32Ptr convert interface{} to *int32.
func ToI32Ptr(source interface{}) *int32 {
	num, err := ToInt32(source)
	if err != nil {
		return nil
	}

	return &num
}

//ToInt16 convert interface{} to int16.
func ToInt16(source interface{}) (int16, error) {
	num, err := ToInt64(source)
	if err != nil {
		return 0, err
	}

	return int16(num), nil
}

//ToI16 convert interface{} to int16.
func ToI16(source interface{}) int16 {
	num, err := ToInt16(source)
	if err != nil {
		return 0
	}

	return num
}

//ToI16Ptr convert interface{} to *int16.
func ToI16Ptr(source interface{}) *int16 {
	num, err := ToInt16(source)
	if err != nil {
		return nil
	}

	return &num
}
