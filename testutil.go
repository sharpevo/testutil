package testutil

import (
	"reflect"
	"testing"
)

func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf("\033[31m"+msg+"\033[39m\n\n", v...)
	}
}

func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("\033[31munexpected error: %s\033[39m\n\n", err.Error())
	}
}

func NotOk(tb testing.TB, err error) {
	tb.Helper()
	if err == nil {
		tb.Fatalf("\033[31mexpected error, got none\033[39m\n\n")
	}
}

func Equal(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\033[31m\n\tEXPECT: %#v\n\tACTUAL: %#v\033[39m\n", exp, act)
	}
}
