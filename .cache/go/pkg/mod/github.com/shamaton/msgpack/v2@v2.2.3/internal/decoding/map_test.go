package decoding

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/shamaton/msgpack/v2/def"
	tu "github.com/shamaton/msgpack/v2/internal/common/testutil"
)

func Test_mapLength(t *testing.T) {
	method := func(d *decoder) func(int, reflect.Kind) (int, int, error) {
		return d.mapLength
	}
	testcases := AsXXXTestCases[int]{
		{
			Name:     "error.code",
			Data:     []byte{},
			Error:    def.ErrTooShortBytes,
			MethodAs: method,
		},
		{
			Name:     "FixMap",
			Data:     []byte{def.FixMap + 3},
			Expected: 3,
			MethodAs: method,
		},
		{
			Name:     "Map16.error",
			Data:     []byte{def.Map16},
			Error:    def.ErrTooShortBytes,
			MethodAs: method,
		},
		{
			Name:     "Map16.ok",
			Data:     []byte{def.Map16, 0xff, 0xff},
			Expected: math.MaxUint16,
			MethodAs: method,
		},
		{
			Name:     "Map32.error",
			Data:     []byte{def.Map32},
			Error:    def.ErrTooShortBytes,
			MethodAs: method,
		},
		{
			Name:     "Map32.ok",
			Data:     []byte{def.Map32, 0xff, 0xff, 0xff, 0xff},
			Expected: math.MaxUint32,
			MethodAs: method,
		},
		{
			Name:     "Unexpected",
			Data:     []byte{def.Nil},
			Error:    def.ErrCanNotDecode,
			MethodAs: method,
		},
	}
	testcases.Run(t)
}

func Test_asFixedMap_StringInt(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asInt",
				Data:           []byte{def.FixStr + 1, 'a', def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.FixStr + 1, 'a', def.PositiveFixIntMin + dv},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[string]int)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[string]int{"a": 1})
	v2 := new(map[string]int8)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[string]int8{"a": 2})
	v3 := new(map[string]int16)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[string]int16{"a": 3})
	v4 := new(map[string]int32)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[string]int32{"a": 4})
	v5 := new(map[string]int64)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[string]int64{"a": 5})
}

func Test_asFixedMap_StringUint(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asUint",
				Data:           []byte{def.FixStr + 1, 'a', def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.FixStr + 1, 'a', def.Uint8, dv},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[string]uint)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[string]uint{"a": 1})
	v2 := new(map[string]uint8)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[string]uint8{"a": 2})
	v3 := new(map[string]uint16)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[string]uint16{"a": 3})
	v4 := new(map[string]uint32)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[string]uint32{"a": 4})
	v5 := new(map[string]uint64)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[string]uint64{"a": 5})
}

func Test_asFixedMap_StringFloat(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asFloat",
				Data:           []byte{def.FixStr + 1, 'a', def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.FixStr + 1, 'a', def.Int16, 0, dv},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[string]float32)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[string]float32{"a": 1})
	v2 := new(map[string]float64)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[string]float64{"a": 2})
}

func Test_asFixedMap_StringBool(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asBool",
				Data:           []byte{def.FixStr + 1, 'a', def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.FixStr + 1, 'a', dv},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[string]bool)
	run(t, v1, def.True)
	tu.EqualMap(t, *v1, map[string]bool{"a": true})
}

func Test_asFixedMap_StringString(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.FixStr + 1, 'a', def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.FixStr + 1, 'a', def.FixStr + 1, dv},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[string]string)
	run(t, v1, 'b')
	tu.EqualMap(t, *v1, map[string]string{"a": "b"})
}

func Test_asFixedMap_IntString(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asInt",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Int8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Int8, dv, def.FixStr + 1, 'b'},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[int]string)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[int]string{1: "b"})
	v2 := new(map[int8]string)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[int8]string{int8(2): "b"})
	v3 := new(map[int16]string)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[int16]string{int16(3): "b"})
	v4 := new(map[int32]string)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[int32]string{int32(4): "b"})
	v5 := new(map[int64]string)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[int64]string{int64(5): "b"})
}

func Test_asFixedMap_IntBool(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asInt",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asBool",
				Data:           []byte{def.Int8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Int8, dv, def.True},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[int]bool)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[int]bool{1: true})
	v2 := new(map[int8]bool)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[int8]bool{int8(2): true})
	v3 := new(map[int16]bool)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[int16]bool{int16(3): true})
	v4 := new(map[int32]bool)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[int32]bool{int32(4): true})
	v5 := new(map[int64]bool)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[int64]bool{int64(5): true})
}

func Test_asFixedMap_UintString(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asUint",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Uint8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Uint8, dv, def.FixStr + 1, 'b'},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[uint]string)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[uint]string{1: "b"})
	v2 := new(map[uint8]string)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[uint8]string{uint8(2): "b"})
	v3 := new(map[uint16]string)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[uint16]string{uint16(3): "b"})
	v4 := new(map[uint32]string)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[uint32]string{uint32(4): "b"})
	v5 := new(map[uint64]string)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[uint64]string{uint64(5): "b"})
}

func Test_asFixedMap_UintBool(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asUint",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asBool",
				Data:           []byte{def.Uint8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Uint8, dv, def.True},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[uint]bool)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[uint]bool{1: true})
	v2 := new(map[uint8]bool)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[uint8]bool{uint8(2): true})
	v3 := new(map[uint16]bool)
	run(t, v3, 3)
	tu.EqualMap(t, *v3, map[uint16]bool{uint16(3): true})
	v4 := new(map[uint32]bool)
	run(t, v4, 4)
	tu.EqualMap(t, *v4, map[uint32]bool{uint32(4): true})
	v5 := new(map[uint64]bool)
	run(t, v5, 5)
	tu.EqualMap(t, *v5, map[uint64]bool{uint64(5): true})
}

func Test_asFixedMap_FloatString(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asFloat",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asString",
				Data:           []byte{def.Uint8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Uint8, dv, def.FixStr + 1, 'b'},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[float32]string)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[float32]string{1: "b"})
	v2 := new(map[float64]string)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[float64]string{2: "b"})
}

func Test_asFixedMap_FloatBool(t *testing.T) {
	run := func(t *testing.T, v any, dv byte) {
		method := func(d *decoder) (int, bool, error) {
			rv := reflect.ValueOf(v)
			return d.asFixedMap(rv.Elem(), 0, 1)
		}

		name := fmt.Sprintf("%T", v)
		testcases := AsXXXTestCases[bool]{
			{
				Name:           name + ".error.asFloat",
				Data:           []byte{def.Str8},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".error.asBool",
				Data:           []byte{def.Uint8, dv, def.Int32},
				Expected:       false,
				Error:          def.ErrCanNotDecode,
				MethodAsCustom: method,
			},
			{
				Name:           name + ".ok",
				Data:           []byte{def.Uint8, dv, def.True},
				Expected:       true,
				MethodAsCustom: method,
			},
		}
		testcases.Run(t)
	}

	v1 := new(map[float32]bool)
	run(t, v1, 1)
	tu.EqualMap(t, *v1, map[float32]bool{1: true})
	v2 := new(map[float64]bool)
	run(t, v2, 2)
	tu.EqualMap(t, *v2, map[float64]bool{2: true})
}
