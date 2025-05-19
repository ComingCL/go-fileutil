// Package tests -----------------------------
// @file      : util_relative_test.go
// @author    : liuChangLe6
// @contact   : liuchangle6@jd.com
// @time      : 2025/5/19 19:59
// -------------------------------------------
package tests

import (
	go_fileutil "github.com/ComingCL/go-fileutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := go_fileutil.ReadFile("./util_relative_test.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", f)
}
