package internal

import (
	"reflect"
	"strconv"
)

func doPrint(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8: // 省略其他长度的类型
		return strconv.FormatInt(v.Int(), 10)
	// ... 为简化起见，省略浮点数和复数分支
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	// 对于引用类型，输出他们的类型以及地址
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	// reflect.Array, reflect.Struct, reflect.Interface
	default:
		return v.Type().String() + " value"
	}
}
