package go_fileutil

import "testing"

func TestReadFile(t *testing.T) {
	f, err := ReadFile("./util_test.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", f)
}
