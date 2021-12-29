package commonutils

import (
	"edge-common-go/pkg/utils"
	"testing"
)

/**
* Description:
*
* @author yang.zhang4
* @date 2020/5/28
 */

func TestPathExists(t *testing.T) {
	_, _ = utils.PathExists("/tmp")
	_, _ = utils.PathExists("/aaa/bbb")
}

func TestGetFileName(t *testing.T) {
	_ = utils.GetFileName()
}

func TestGetLineNum(t *testing.T) {
	_ = utils.GetLineNum()
}

func TestGetFuncName(t *testing.T) {
	_ = utils.GetFuncName()
}

func TestFloatIsEqual(t *testing.T) {
	var isEqual bool
	var double1 = 0.0001
	var double2 = 0.00011
	isEqual = utils.FloatIsEqual(double1, double2)
	if isEqual {
		t.Fatalf("%f == %f", double1, double2)
	}

	double1 = 1.1
	double2 = 1.1000009
	isEqual = utils.FloatIsEqual(double1, double2)
	if !isEqual {
		t.Fatalf("%f != %f", double1, double2)
	}
}

func TestFloatIsEqualX(t *testing.T) {
	var isEqual bool
	var double1 = 0.0001
	var double2 = 0.00011
	var x = 0.000001
	isEqual = utils.FloatIsEqualX(double1, double2, x)
	if isEqual {
		t.Fatalf("%f == %f", double1, double2)
	}

	x = 0.00002
	isEqual = utils.FloatIsEqualX(double1, double2, x)
	if !isEqual {
		t.Fatalf("%f != %f", double1, double2)
	}
}

func TestNewStringPtr(t *testing.T) {
	_ = utils.NewStringPtr("aaa")
}

func TestGetValueFromStringPtr(t *testing.T) {
	var stringValue = "aaa"
	var valueCopy string
	valueCopy = utils.GetValueFromStringPtr(&stringValue, "bbb")

	if valueCopy != stringValue {
		t.Fatalf("valueCopy != stringValue")
	}

	valueCopy = utils.GetValueFromStringPtr(nil, "bbb")
	if valueCopy != "bbb" {
		t.Fatalf("valueCopy != bbb")
	}
}

func TestCopyStringPtrNotNil(t *testing.T) {
	var stringValue = "aaa"
	var valueCopy *string
	valueCopy = utils.CopyStringPtrNotNil(&stringValue, "bbb")
	if *valueCopy != stringValue {
		t.Fatalf("*valueCopy != stringValue")
	}

	valueCopy = utils.CopyStringPtrNotNil(nil, "bbb")
	if *valueCopy != "bbb" {
		t.Fatalf("*valueCopy != bbb")
	}
}

func TestNewBoolPtr(t *testing.T) {
	_ = utils.NewBoolPtr(true)
}

func TestGetValueFromBoolPtr(t *testing.T) {
	var srcValue = true
	var valueCopy bool
	valueCopy = utils.GetValueFromBoolPtr(&srcValue, false)

	if valueCopy != srcValue {
		t.Fatalf("valueCopy != srcValue")
	}

	valueCopy = utils.GetValueFromBoolPtr(nil, true)
	if valueCopy != true {
		t.Fatalf("valueCopy != true")
	}
}

func TestCopyBoolPtrNotNil(t *testing.T) {
	var srcValue = true
	var valueCopy *bool
	valueCopy = utils.CopyBoolPtrNotNil(&srcValue, false)
	if *valueCopy != srcValue {
		t.Fatalf("*valueCopy != srcValue")
	}

	valueCopy = utils.CopyBoolPtrNotNil(nil, true)
	if *valueCopy != true {
		t.Fatalf("*valueCopy != true")
	}
}

func TestNewInt32Ptr(t *testing.T) {
	_ = utils.NewInt32Ptr(30)
}

func TestGetValueFromInt32Ptr(t *testing.T) {
	var value int32 = 10
	var valueCopy int32
	valueCopy = utils.GetValueFromInt32Ptr(&value, 20)

	if valueCopy != value {
		t.Fatalf("valueCopy != value")
	}

	valueCopy = utils.GetValueFromInt32Ptr(nil, 20)
	if valueCopy != 20 {
		t.Fatalf("valueCopy != 20")
	}
}

func TestCopyInt32PtrNotNil(t *testing.T) {
	var value int32 = 10
	var valueCopy *int32
	valueCopy = utils.CopyInt32PtrNotNil(&value, 20)
	if *valueCopy != value {
		t.Fatalf("*valueCopy != value")
	}

	valueCopy = utils.CopyInt32PtrNotNil(nil, 20)
	if *valueCopy != 20 {
		t.Fatalf("valueCopy != 20")
	}
}

func TestNewInt64Ptr(t *testing.T) {
	_ = utils.NewInt64Ptr(30)
}

func TestGetValueFromInt64Ptr(t *testing.T) {
	var value int64 = 10
	var valueCopy int64
	valueCopy = utils.GetValueFromInt64Ptr(&value, 20)

	if valueCopy != value {
		t.Fatalf("valueCopy != value")
	}

	valueCopy = utils.GetValueFromInt64Ptr(nil, 20)
	if valueCopy != 20 {
		t.Fatalf("valueCopy != 20")
	}
}

func TestCopyInt64PtrNotNil(t *testing.T) {
	var value int64 = 10
	var valueCopy *int64
	valueCopy = utils.CopyInt64PtrNotNil(&value, 20)
	if *valueCopy != value {
		t.Fatalf("*valueCopy != value")
	}

	valueCopy = utils.CopyInt64PtrNotNil(nil, 20)
	if *valueCopy != 20 {
		t.Fatalf("valueCopy != 20")
	}
}

func TestStringIsInArray(t *testing.T) {
	var str = "aaa"
	var strArray = []string{"aaa", "bbb"}
	var result = utils.StringIsInArray(strArray, str)
	if !result {
		t.Fatalf("str is not in strArray")
	}

	str = "ccc"
	result = utils.StringIsInArray(strArray, str)
	if result {
		t.Fatalf("str is in strArray")
	}
}

func TestCombineStringArray(t *testing.T) {
	var strArray1 = []string{"aaa", "bbb"}
	var strArray2 = []string{"bbb", "ccc"}
	var strArray3 = []string{"ddd", "eee"}
	var strArrayExpect = []string{"aaa", "bbb", "ccc", "ddd", "eee"}

	var strArrayResult = utils.CombineStringArray(strArray1, strArray2, strArray3)

	if !stringSliceIsEqual(strArrayResult, strArrayExpect) {
		t.Fatalf("strArrayResult != strArrayExpect")
	}
}

func TestGetBaseDir(t *testing.T) {
	_ = utils.GetBaseDir()
}

func TestDeepCopyMap(t *testing.T) {
	var mapSrc = map[string]interface{}{
		"aaa": "bbb",
		"ccc": "ddd",
	}
	var mapDst = utils.DeepCopyMap(mapSrc)

	if mapDst["aaa"] != mapSrc["aaa"] {
		t.Fatalf(`mapDst["aaa"] != mapSrc["aaa"]`)
	}

	if mapDst["ccc"] != mapSrc["ccc"] {
		t.Fatalf(`mapDst["ccc"] != mapSrc["ccc"]`)
	}
}

func stringSliceIsEqual(slice1 []string, slice2 []string) bool {
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
