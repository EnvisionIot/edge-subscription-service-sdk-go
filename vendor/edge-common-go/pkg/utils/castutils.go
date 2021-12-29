package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"html/template"
	"math"
	"reflect"
	"strconv"
)

/**
* Description:
*
* @author yang.zhang4
* @date 2020/7/23
 */

//interface转为float64，支持数组
func ToFloat64E(i interface{}) (float64, error) {
	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64, i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case []interface{}:
		if len(s) <= 0 {
			return 0, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64, i, i)
		}
		return cast.ToFloat64E(s[0])
	default:
		return 0, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64, i, i)
	}
}

//interface转为string，支持数组
func ToStringE(i interface{}) (string, error) {
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatInt(int64(s), 10), nil
	case uint64:
		return strconv.FormatInt(int64(s), 10), nil
	case uint32:
		return strconv.FormatInt(int64(s), 10), nil
	case uint16:
		return strconv.FormatInt(int64(s), 10), nil
	case uint8:
		return strconv.FormatInt(int64(s), 10), nil
	case []byte:
		return string(s), nil
	case []interface{}:
		if len(s) <= 0 {
			return "", fmt.Errorf(UNABLE_TO_CAST_TYPE+STRING, i, i)
		}
		return cast.ToStringE(s[0])
	case map[string]interface{}:
		var err error
		var byteArray []byte
		byteArray, err = json.Marshal(s)
		if err != nil {
			return "", err
		}
		return string(byteArray), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf(UNABLE_TO_CAST_TYPE+STRING, i, i)
	}
}

//从float64转为int64，当超过int32最大范围时，用最大范围的值表示
func GetInt64FromFloat64WithInt32Limit(val float64) int64 {
	var valueInt64Temp int64
	if val >= math.MaxInt32 {
		valueInt64Temp = math.MaxInt32
	} else if val <= math.MinInt32 {
		valueInt64Temp = math.MinInt32
	} else {
		valueInt64Temp = int64(val)
	}

	return valueInt64Temp
}

//从float64转为int64，当超过int64最大范围时，用最大范围的值表示
func GetInt64FromFloat64WithInt64Limit(val float64) int64 {
	var valueInt64Temp int64
	if val >= math.MaxInt64 {
		valueInt64Temp = math.MaxInt64
	} else if val <= math.MinInt64 {
		valueInt64Temp = math.MinInt64
	} else {
		valueInt64Temp = int64(val)
	}

	return valueInt64Temp
}

//从float64转为float32，当超过float32最大范围时，用最大范围的值表示
func GetFloat64FromFloat64WithFloat32Limit(val float64) float64 {
	var valueFloat64Temp float64
	if val >= math.MaxFloat32 {
		valueFloat64Temp = math.MaxFloat32
	} else if val <= (-1)*math.MaxFloat32 {
		valueFloat64Temp = (-1) * math.MaxFloat32
	} else {
		valueFloat64Temp = val
	}

	return valueFloat64Temp
}

// ToInt32SliceE casts an interface to a []int32 type.
func ToInt32SliceE(i interface{}) ([]int32, error) {
	if i == nil {
		return []int32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT32_SLICE, i, i)
	}

	switch v := i.(type) {
	case []int32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int32, s.Len())
		for j := 0; j < s.Len(); j++ {
			valueFloat64Temp, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return []int32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT32_SLICE, i, i)
			}

			var valueInt64Temp = GetInt64FromFloat64WithInt32Limit(valueFloat64Temp)
			a[j] = int32(valueInt64Temp)
		}
		return a, nil
	default:
		valueFloat64Temp, err := cast.ToFloat64E(i)
		if err != nil {
			return []int32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT32_SLICE, i, i)
		}
		var valueInt64Temp = GetInt64FromFloat64WithInt32Limit(valueFloat64Temp)
		return []int32{int32(valueInt64Temp)}, nil
	}
}

// ToInt64SliceE casts an interface to a []int64 type.
func ToInt64SliceE(i interface{}) ([]int64, error) {
	if i == nil {
		return []int64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT64_SLICE, i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			valueFloat64Temp, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return []int64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT64_SLICE, i, i)
			}

			var valueInt64Temp = GetInt64FromFloat64WithInt64Limit(valueFloat64Temp)
			a[j] = valueInt64Temp
		}
		return a, nil
	default:
		valueFloat64Temp, err := cast.ToFloat64E(i)
		if err != nil {
			return []int64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+INT64_SLICE, i, i)
		}
		var valueInt64Temp = GetInt64FromFloat64WithInt64Limit(valueFloat64Temp)
		return []int64{valueInt64Temp}, nil
	}
}

// ToFloat32SliceE casts an interface to a []float32 type.
func ToFloat32SliceE(i interface{}) ([]float32, error) {
	if i == nil {
		return []float32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT32_SLICE, i, i)
	}

	switch v := i.(type) {
	case []float32:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]float32, s.Len())
		for j := 0; j < s.Len(); j++ {
			valueFloat64Temp, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return []float32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT32_SLICE, i, i)
			}

			valueFloat64Temp = GetFloat64FromFloat64WithFloat32Limit(valueFloat64Temp)
			a[j] = float32(valueFloat64Temp)
		}
		return a, nil
	default:
		valueFloat64Temp, err := cast.ToFloat64E(i)
		if err != nil {
			return []float32{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT32_SLICE, i, i)
		}
		valueFloat64Temp = GetFloat64FromFloat64WithFloat32Limit(valueFloat64Temp)
		return []float32{float32(valueFloat64Temp)}, nil
	}
}

// ToFloat64SliceE casts an interface to a []float64 type.
func ToFloat64SliceE(i interface{}) ([]float64, error) {
	if i == nil {
		return []float64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64_SLICE, i, i)
	}

	switch v := i.(type) {
	case []float64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]float64, s.Len())
		for j := 0; j < s.Len(); j++ {
			valueFloat64Temp, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return []float64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64_SLICE, i, i)
			}

			a[j] = valueFloat64Temp
		}
		return a, nil
	default:
		valueFloat64Temp, err := cast.ToFloat64E(i)
		if err != nil {
			return []float64{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+FLOAT64_SLICE, i, i)
		}
		return []float64{valueFloat64Temp}, nil
	}
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE(i interface{}) ([]string, error) {
	if i == nil {
		return []string{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+STRING_SLICE, i, i)
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]string, s.Len())
		for j := 0; j < s.Len(); j++ {
			valueStringTemp, err := cast.ToStringE(s.Index(j).Interface())
			if err != nil {
				return []string{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+STRING_SLICE, i, i)
			}

			a[j] = valueStringTemp
		}
		return a, nil
	default:
		valueStringTemp, err := cast.ToStringE(i)
		if err != nil {
			return []string{}, fmt.Errorf(UNABLE_TO_CAST_TYPE+STRING_SLICE, i, i)
		}
		return []string{valueStringTemp}, nil
	}
}
