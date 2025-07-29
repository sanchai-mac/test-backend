package testutil

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("error is not nil: %v", err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal(err)
	}
}

func IsError(t *testing.T, actual, expected error) {
	t.Helper()
	if !errors.Is(actual, expected) {
		t.Fatalf("not equal. actual: %v, expected: %v", actual, expected)
	}
}

func ErrorContains(t *testing.T, err error, errStr string) {
	t.Helper()
	if err == nil {
		t.Fatal("error should occur")
	}
	if !strings.Contains(err.Error(), errStr) {
		t.Fatalf("error does not contain '%s'. err: %v", errStr, err)
	}
}

func Equal[T any](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("not equal. actual: %v, expected: %v", actual, expected)
	}
}

func EqualSlice[T comparable](t *testing.T, actual, expected []T) {
	t.Helper()
	if len(actual) != len(expected) {
		switch a := any(actual).(type) {
		case []byte:
			e := any(expected).([]byte)
			t.Fatalf("diffrent length. actual: [% 02x], expected: [% 02x]", a, e)
		default:
			t.Fatalf("diffrent length. actual: %v, expected: %v", actual, expected)
		}
	}
	for i := range actual {
		if !reflect.DeepEqual(actual[i], expected[i]) {
			switch a := any(actual).(type) {
			case []byte:
				e := any(expected).([]byte)
				t.Fatalf("not equal. actual: [% 02x], expected: [% 02x]", a, e)
			default:
				t.Fatalf("not equal. actual: %v, expected: %v", actual, expected)
			}
		}
	}
}

func EqualMap[K comparable, V comparable](t *testing.T, actual, expected map[K]V) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("diffrent length. actual: %v, expected: %v", actual, expected)
	}
	for k, v1 := range actual {
		if v2, ok := expected[k]; !ok || v1 != v2 {
			t.Fatalf("not equal. actual: %v, expected: %v", actual, expected)
		}
	}
}

type Equaler[T any] interface {
	Equal(other T) bool
}

func EqualEqualer[T Equaler[T]](t *testing.T, actual, expected T) {
	t.Helper()
	if !actual.Equal(expected) {
		t.Fatalf("not equal. actual: %v, expected: %v", actual, expected)
	}
}
