package test

import (
	"path/filepath"
	"testing"

	"os"

	"github.com/ariefdarmawan/gopipe"
	"github.com/eaciit/toolkit"
)

var (
	wd, _    = os.Getwd()
	codePath = filepath.Join(wd, "..", "code")
)

func TestCmd(t *testing.T) {
	name := "Go Pipe"
	cmd := gopipe.NewCommand(gopipe.CommandDirect, "fmt.Println(\"Hello %s\", $name)")
	cmd.AddLib("fmt")
	result := cmd.Exec(toolkit.M{}.Set("name", name))
	checkResult(t, result, "Hello "+name)
}

func TestFile(t *testing.T) {
	name := "Go Pipe"
	cmd := gopipe.NewCommand(gopipe.CommandFile, filepath.Join(codePath, "sample2"))
	result := cmd.Exec(toolkit.M{}.Set("name", name))
	checkResult(t, result, "Hello "+name)
}

func checkResult(t *testing.T, r *toolkit.Result, want interface{}) {
	if r.Status != toolkit.Status_OK {
		t.Errorf("Fail to get result: %s", r.Message)
	}

	check(t, want, r.Data)
}

func check(t *testing.T, want interface{}, got interface{}) {
	if toolkit.IsNilOrEmpty(want) {
		t.Error("Expected value is nil")
	}

	if toolkit.IsNilOrEmpty(got) {
		t.Error("Output value is nil")
	}

	if want != got {
		t.Errorf("Want \"%v\" got \"%v\"", want, got)
	}
}
