package castutils

import (
	"edge-common-go/pkg/utils"
	"github.com/spf13/cast"
	"math"
	"testing"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2020/12/3
 */

func TestToFloat64E(t *testing.T) {
	var stringValue = "5.5"
	var float64Value float64
	var valueExpect = 5.5
	var err error
	float64Value, err = utils.ToFloat64E(stringValue)
	if err != nil {
		t.Fatalf("utils.ToFloat64E error, %s", err)
	}

	if !utils.FloatIsEqual(float64Value, valueExpect) {
		t.Fatalf("float64Value(%f) != valueExpect(%f)", float64Value, valueExpect)
	}

	var sliceValue = []interface{}{"6.6", "7.7"}
	valueExpect = 6.6
	float64Value, err = utils.ToFloat64E(sliceValue)
	if err != nil {
		t.Fatalf("utils.ToFloat64E error, %s", err)
	}

	if !utils.FloatIsEqual(float64Value, valueExpect) {
		t.Fatalf("float64Value(%f) != valueExpect(%f)", float64Value, valueExpect)
	}
}

func TestToStringE(t *testing.T) {
	var stringValue string
	var int64Value = 100
	var valueExpect = "100"
	var err error
	stringValue, err = utils.ToStringE(int64Value)
	if err != nil {
		t.Fatalf("utils.ToStringE error, %s", err)
	}

	if stringValue != valueExpect {
		t.Fatalf("stringValue(%s) != valueExpect(%s)", stringValue, valueExpect)
	}

	var sliceValue = []interface{}{200, 300}
	valueExpect = "200"
	stringValue, err = utils.ToStringE(sliceValue)
	if err != nil {
		t.Fatalf("utils.ToStringE error, %s", err)
	}

	if stringValue != valueExpect {
		t.Fatalf("stringValue(%s) != valueExpect(%s)", stringValue, valueExpect)
	}
}

func TestGetInt64FromFloat64WithInt32Limit(t *testing.T) {
	var valueIn float64 = math.MaxInt32 + 1
	var valueRet = utils.GetInt64FromFloat64WithInt32Limit(valueIn)
	if valueRet != math.MaxInt32 {
		t.Fatalf("valueRet(%d) != math.MaxInt32", valueRet)
	}

	valueIn = math.MinInt32 - 1
	valueRet = utils.GetInt64FromFloat64WithInt32Limit(valueIn)
	if valueRet != math.MinInt32 {
		t.Fatalf("valueRet(%d) != math.MinInt32", valueRet)
	}

	valueIn = 10
	valueRet = utils.GetInt64FromFloat64WithInt32Limit(valueIn)
	if float64(valueRet) != valueIn {
		t.Fatalf("valueRet(%d) != valueIn(%f)", valueRet, valueIn)
	}
}

func TestGetInt64FromFloat64WithInt64Limit(t *testing.T) {
	var valueIn float64 = math.MaxInt64 * 2
	var valueRet = utils.GetInt64FromFloat64WithInt64Limit(valueIn)
	if valueRet != math.MaxInt64 {
		t.Fatalf("valueRet(%d) != math.MaxInt64", valueRet)
	}

	valueIn = math.MinInt64 * 2
	valueRet = utils.GetInt64FromFloat64WithInt64Limit(valueIn)
	if valueRet != math.MinInt64 {
		t.Fatalf("valueRet(%d) != math.MinInt64", valueRet)
	}

	valueIn = 10
	valueRet = utils.GetInt64FromFloat64WithInt64Limit(valueIn)
	if float64(valueRet) != valueIn {
		t.Fatalf("valueRet(%d) != valueIn(%f)", valueRet, valueIn)
	}
}

func TestGetFloat64FromFloat64WithFloat32Limit(t *testing.T) {
	var valueIn = math.MaxFloat32 + 1
	var valueRet = utils.GetFloat64FromFloat64WithFloat32Limit(valueIn)
	if valueRet != math.MaxFloat32 {
		t.Fatalf("valueRet(%f) != math.MaxFloat32", valueRet)
	}

	valueIn = (-1)*math.MaxFloat32 - 1
	valueRet = utils.GetFloat64FromFloat64WithFloat32Limit(valueIn)
	if valueRet != (-1)*math.MaxFloat32 {
		t.Fatalf("valueRet(%f) != math.MaxFloat32", valueRet)
	}

	valueIn = 10
	valueRet = utils.GetFloat64FromFloat64WithFloat32Limit(valueIn)
	if valueRet != valueIn {
		t.Fatalf("valueRet(%f) != valueIn(%f)", valueRet, valueIn)
	}
}

func TestToInt32SliceE(t *testing.T) {
	var valueRet []int32
	var err error

	valueRet, err = utils.ToInt32SliceE(nil)
	if err == nil {
		t.Fatalf("err is nil when input is nil")
	}

	var valueInInt32Slice = []int32{1, 2, 3}
	var valueExpect = []int32{1, 2, 3}
	valueRet, err = utils.ToInt32SliceE(valueInInt32Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	var isEqual bool
	isEqual = int32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64Slice = []float64{1.1, 2.2, 6.6}
	valueExpect = []int32{1, 2, 6}
	valueRet, err = utils.ToInt32SliceE(valueInFloat64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInStringSlice = []string{"1.1", "2.2", "6.6"}
	valueExpect = []int32{1, 2, 6}
	valueRet, err = utils.ToInt32SliceE(valueInStringSlice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64 = 7.7
	valueExpect = []int32{7}
	valueRet, err = utils.ToInt32SliceE(valueInFloat64)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}
}

func TestToInt64SliceE(t *testing.T) {
	var valueRet []int64
	var err error

	valueRet, err = utils.ToInt64SliceE(nil)
	if err == nil {
		t.Fatalf("err is nil when input is nil")
	}

	var valueInInt64Slice = []int64{1, 2, 3}
	var valueExpect = []int64{1, 2, 3}
	valueRet, err = utils.ToInt64SliceE(valueInInt64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	var isEqual bool
	isEqual = int64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64Slice = []float64{1.1, 2.2, 6.6}
	valueExpect = []int64{1, 2, 6}
	valueRet, err = utils.ToInt64SliceE(valueInFloat64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInStringSlice = []string{"1.1", "2.2", "6.6"}
	valueExpect = []int64{1, 2, 6}
	valueRet, err = utils.ToInt64SliceE(valueInStringSlice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64 = 7.7
	valueExpect = []int64{7}
	valueRet, err = utils.ToInt64SliceE(valueInFloat64)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = int64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}
}

func TestToFloat32SliceE(t *testing.T) {
	var valueRet []float32
	var err error

	valueRet, err = utils.ToFloat32SliceE(nil)
	if err == nil {
		t.Fatalf("err is nil when input is nil")
	}

	var valueInInt32Slice = []int32{1, 2, 3}
	var valueExpect = []float32{1, 2, 3}
	valueRet, err = utils.ToFloat32SliceE(valueInInt32Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	var isEqual bool
	isEqual = float32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64Slice = []float64{1.1, 2.2, 6.6}
	valueExpect = []float32{1.1, 2.2, 6.6}
	valueRet, err = utils.ToFloat32SliceE(valueInFloat64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInStringSlice = []string{"1.1", "2.2", "6.6"}
	valueExpect = []float32{1.1, 2.2, 6.6}
	valueRet, err = utils.ToFloat32SliceE(valueInStringSlice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64 = 7.7
	valueExpect = []float32{7.7}
	valueRet, err = utils.ToFloat32SliceE(valueInFloat64)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float32SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}
}

func TestToFloat64SliceE(t *testing.T) {
	var valueRet []float64
	var err error

	valueRet, err = utils.ToFloat64SliceE(nil)
	if err == nil {
		t.Fatalf("err is nil when input is nil")
	}

	var valueInInt32Slice = []int32{1, 2, 3}
	var valueExpect = []float64{1, 2, 3}
	valueRet, err = utils.ToFloat64SliceE(valueInInt32Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	var isEqual bool
	isEqual = float64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64Slice = []float64{1.1, 2.2, 6.6}
	valueExpect = []float64{1.1, 2.2, 6.6}
	valueRet, err = utils.ToFloat64SliceE(valueInFloat64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInStringSlice = []string{"1.1", "2.2", "6.6"}
	valueExpect = []float64{1.1, 2.2, 6.6}
	valueRet, err = utils.ToFloat64SliceE(valueInStringSlice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64 = 7.7
	valueExpect = []float64{7.7}
	valueRet, err = utils.ToFloat64SliceE(valueInFloat64)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = float64SliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}
}

func TestToStringSliceE(t *testing.T) {
	var valueRet []string
	var err error

	valueRet, err = utils.ToStringSliceE(nil)
	if err == nil {
		t.Fatalf("err is nil when input is nil")
	}

	var valueInInt32Slice = []int32{1, 2, 3}
	var valueExpect = []string{"1", "2", "3"}
	valueRet, err = utils.ToStringSliceE(valueInInt32Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	var isEqual bool
	isEqual = stringSliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64Slice = []float64{1.1, 2.2, 6.6}
	valueExpect = []string{"1.1", "2.2", "6.6"}
	valueRet, err = utils.ToStringSliceE(valueInFloat64Slice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = stringSliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInStringSlice = []string{"1.1", "2.2", "6.6"}
	valueExpect = []string{"1.1", "2.2", "6.6"}
	valueRet, err = utils.ToStringSliceE(valueInStringSlice)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = stringSliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}

	var valueInFloat64 = 7.7
	valueExpect = []string{"7.7"}
	valueRet, err = utils.ToStringSliceE(valueInFloat64)
	if err != nil {
		t.Fatalf("err=%s", err)
	}
	isEqual = stringSliceIsEqual(valueRet, valueExpect)
	if !isEqual {
		t.Fatalf("valueRet=%+v != valueExpect=%+v", valueRet, valueExpect)
	}
}

func int32SliceIsEqual(slice1 []int32, slice2 []int32) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for ii := 0; ii < len(slice1); ii++ {
		if slice1[ii] != slice2[ii] {
			return false
		}
	}

	return true
}

func int64SliceIsEqual(slice1 []int64, slice2 []int64) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for ii := 0; ii < len(slice1); ii++ {
		if slice1[ii] != slice2[ii] {
			return false
		}
	}

	return true
}

func float32SliceIsEqual(slice1 []float32, slice2 []float32) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for ii := 0; ii < len(slice1); ii++ {
		if !utils.FloatIsEqual(float64(slice1[ii]), float64(slice2[ii])) {
			return false
		}
	}

	return true
}

func float64SliceIsEqual(slice1 []float64, slice2 []float64) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for ii := 0; ii < len(slice1); ii++ {
		if !utils.FloatIsEqual(slice1[ii], slice2[ii]) {
			return false
		}
	}

	return true
}

func stringSliceIsEqual(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for ii := 0; ii < len(slice1); ii++ {
		var float64Temp1 = cast.ToFloat64(slice1[ii])
		var float64Temp2 = cast.ToFloat64(slice2[ii])
		if !utils.FloatIsEqual(float64Temp1, float64Temp2) {
			return false
		}
	}

	return true
}
